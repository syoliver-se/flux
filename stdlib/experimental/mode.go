package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	modeSignature := runtime.MustLookupBuiltinType("experimental", "mode")
	runtime.RegisterPackageValue("experimental", "mode", flux.MustValue(flux.FunctionValue("mode", universe.CreateModeOpSpec, modeSignature)))
}
