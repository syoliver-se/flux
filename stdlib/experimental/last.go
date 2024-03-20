package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	lastSignature := runtime.MustLookupBuiltinType("experimental", "last")
	runtime.RegisterPackageValue("experimental", "last", flux.MustValue(flux.FunctionValue("last", universe.CreateLastOpSpec, lastSignature)))
}
