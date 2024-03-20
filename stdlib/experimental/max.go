package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	maxSignature := runtime.MustLookupBuiltinType("experimental", "max")
	runtime.RegisterPackageValue("experimental", "max", flux.MustValue(flux.FunctionValue("max", universe.CreateMaxOpSpec, maxSignature)))
}
