package execute

import (
	"context"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/compiler"
	"github.com/syoliver-se/flux/execute/table"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/semantic"
	"github.com/syoliver-se/flux/values"
)

type VectorMapFn struct {
	dynamicFn
}

func NewVectorMapFn(fn *semantic.FunctionExpression, scope compiler.Scope) *VectorMapFn {
	return &VectorMapFn{
		dynamicFn: newDynamicFn(fn, scope),
	}
}

func (f *VectorMapFn) Prepare(ctx context.Context, cols []flux.ColMeta) (*VectorMapPreparedFn, error) {
	fn, err := f.prepare(ctx, cols, nil, true)
	if err != nil {
		return nil, err
	} else if k := fn.returnType().Nature(); k != semantic.Object {
		return nil, errors.Newf(codes.Invalid, "map function must return an object, got %s", k.String())
	}
	return &VectorMapPreparedFn{
		vectorFn: vectorFn{preparedFn: fn},
	}, nil
}

type VectorMapPreparedFn struct {
	vectorFn
}

func (f *VectorMapPreparedFn) Type() semantic.MonoType {
	return f.fn.Type()
}

type vectorFn struct {
	preparedFn
}

func (f *vectorFn) Eval(ctx context.Context, chunk table.Chunk) (values.Object, error) {
	for j, col := range chunk.Cols() {
		arr := chunk.Values(j)
		arr.Retain()
		v := values.NewVectorValue(arr, flux.SemanticType(col.Type))
		f.arg0.Set(col.Label, v)
	}
	defer f.arg0.Release()

	res, err := f.fn.Eval(ctx, f.args)
	if err != nil {
		return nil, err
	}
	return res.Object(), nil
}
