package experimental

import (
	"context"

	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/runtime"

	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/semantic"
	"github.com/syoliver-se/flux/values"
)

func init() {
	runtime.RegisterPackageValue("experimental", "objectKeys", values.NewFunction(
		"objectKeys",
		runtime.MustLookupBuiltinType("experimental", "objectKeys"),
		func(ctx context.Context, args values.Object) (values.Value, error) {
			o, ok := args.Get("o")
			if !ok {
				return nil, errors.New(codes.Invalid, "missing parameter \"o\"")
			}
			if o.Type().Nature() != semantic.Object {
				return nil, errors.New(codes.Invalid, "parameter \"o\" is not an object")
			}
			obj := o.Object()
			keys := make([]values.Value, 0, obj.Len())
			obj.Range(func(name string, _ values.Value) {
				keys = append(keys, values.NewString(name))
			})
			return values.NewArrayWithBacking(semantic.NewArrayType(semantic.BasicString), keys), nil
		},
		false,
	))
}
