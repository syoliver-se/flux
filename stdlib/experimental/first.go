package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	firstSignature := runtime.MustLookupBuiltinType("experimental", "first")
	runtime.RegisterPackageValue("experimental", "first", flux.MustValue(flux.FunctionValue("first", universe.CreateFirstOpSpec, firstSignature)))
}
