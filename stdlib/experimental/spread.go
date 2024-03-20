package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	spreadSignature := runtime.MustLookupBuiltinType("experimental", "spread")
	runtime.RegisterPackageValue("experimental", "spread", flux.MustValue(flux.FunctionValue("spread", universe.CreateSpreadOpSpec, spreadSignature)))
}
