package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	distinctSignature := runtime.MustLookupBuiltinType("experimental", "distinct")
	runtime.RegisterPackageValue("experimental", "distinct", flux.MustValue(flux.FunctionValue("distinct", universe.CreateDistinctOpSpec, distinctSignature)))
}
