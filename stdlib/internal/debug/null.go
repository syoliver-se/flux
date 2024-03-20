package debug

import (
	"context"

	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/interpreter"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/semantic"
	"github.com/syoliver-se/flux/values"
)

func init() {
	mt := runtime.MustLookupBuiltinType("internal/debug", "null")
	runtime.RegisterPackageValue("internal/debug", "null",
		values.NewFunction("null", mt, func(ctx context.Context, args values.Object) (values.Value, error) {
			return interpreter.DoFunctionCallContext(Null, ctx, args)
		}, false),
	)
}

func Null(ctx context.Context, args interpreter.Arguments) (values.Value, error) {
	typ, ok, err := args.GetString("type")
	if err != nil {
		return nil, err
	}

	if !ok {
		return values.Null, nil
	}

	var semanticType semantic.MonoType
	switch typ {
	case "string":
		semanticType = semantic.BasicString
	case "bytes":
		semanticType = semantic.BasicBytes
	case "int":
		semanticType = semantic.BasicInt
	case "uint":
		semanticType = semantic.BasicUint
	case "float":
		semanticType = semantic.BasicFloat
	case "bool":
		semanticType = semantic.BasicBool
	case "time":
		semanticType = semantic.BasicTime
	case "duration":
		semanticType = semantic.BasicDuration
	case "regexp":
		semanticType = semantic.BasicRegexp
	default:
		return nil, errors.Newf(codes.Invalid, "Invalid null type `%s`", typ)
	}

	return values.NewNull(semanticType), nil
}
