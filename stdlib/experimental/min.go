package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	minSignature := runtime.MustLookupBuiltinType("experimental", "min")
	runtime.RegisterPackageValue("experimental", "min", flux.MustValue(flux.FunctionValue("min", universe.CreateMinOpSpec, minSignature)))
}
