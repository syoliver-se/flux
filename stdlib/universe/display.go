package universe

import (
	"context"

	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/values"
)

func init() {
	runtime.RegisterPackageValue("universe", "display", values.NewFunction(
		"display",
		runtime.MustLookupBuiltinType("universe", "display"),
		func(ctx context.Context, args values.Object) (values.Value, error) {
			v, ok := args.Get("v")
			if !ok {
				return nil, errors.New(codes.Invalid, "missing argument v")
			}
			return values.NewString(values.DisplayString(v)), nil
		},
		false,
	))
}
