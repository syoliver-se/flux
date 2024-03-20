package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	windowSignature := runtime.MustLookupBuiltinType("experimental", "_window")
	runtime.RegisterPackageValue("experimental", "_window", flux.MustValue(flux.FunctionValue("window", universe.CreateWindowOpSpec, windowSignature)))
}
