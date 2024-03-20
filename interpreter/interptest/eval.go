package interptest

import (
	"context"

	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/interpreter"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/values"
)

func Eval(ctx context.Context, itrp *interpreter.Interpreter, scope values.Scope, importer interpreter.Importer, src string) ([]interpreter.SideEffect, error) {
	node, err := runtime.AnalyzeSource(ctx, src)
	if err != nil {
		return nil, errors.Wrap(err, codes.Inherit, "could not analyze program")
	}
	return itrp.Eval(ctx, node, scope, importer)
}
