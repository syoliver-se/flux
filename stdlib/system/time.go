package system

import (
	"context"
	"time"

	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/semantic"
	"github.com/syoliver-se/flux/values"
)

var systemTimeFuncName = "time"

func init() {
	runtime.RegisterPackageValue("system", systemTimeFuncName, values.NewFunction(
		systemTimeFuncName,
		semantic.NewFunctionType(semantic.BasicTime, nil),
		func(ctx context.Context, args values.Object) (values.Value, error) {
			return values.NewTime(values.ConvertTime(time.Now().UTC())), nil
		},
		false,
	))
}
