package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	countSignature := runtime.MustLookupBuiltinType("experimental", "count")
	runtime.RegisterPackageValue("experimental", "count", flux.MustValue(flux.FunctionValue("count", universe.CreateCountOpSpec, countSignature)))
}
