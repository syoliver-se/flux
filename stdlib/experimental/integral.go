package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	integralSignature := runtime.MustLookupBuiltinType("experimental", "integral")
	runtime.RegisterPackageValue("experimental", "integral", flux.MustValue(flux.FunctionValue("integral", universe.CreateIntegralOpSpec, integralSignature)))
}
