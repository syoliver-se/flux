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

func getStringArgument(args values.Object, name string) (string, error) {
	v, ok := args.Get(name)
	if !ok {
		return "", errors.Newf(codes.Invalid, "missing argument %s", name)
	}
	if v.Type().Nature() != semantic.String {
		return "", errors.Newf(codes.Invalid, "cannot convert argument of type %v to string", v.Type().Nature())
	}
	return v.Str(), nil
}

func init() {
	name := "getOption"
	runtime.RegisterPackageValue("internal/debug", name, values.NewFunction(
		name,
		runtime.MustLookupBuiltinType("internal/debug", name),
		func(ctx context.Context, args values.Object) (values.Value, error) {
			pkg, err := getStringArgument(args, "pkg")
			if err != nil {
				return nil, err
			}

			name, err := getStringArgument(args, "name")
			if err != nil {
				return nil, err
			}

			v, ok := interpreter.GetOption(ctx, pkg, name)
			if !ok {
				return nil, errors.Newf(codes.Invalid, "option does not exist")
			}

			return v, nil
		}, false,
	))
}
