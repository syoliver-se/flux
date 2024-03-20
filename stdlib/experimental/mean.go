package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	meanSignature := runtime.MustLookupBuiltinType("experimental", "mean")
	runtime.RegisterPackageValue("experimental", "mean", flux.MustValue(flux.FunctionValue("mean", universe.CreateMeanOpSpec, meanSignature)))
}
