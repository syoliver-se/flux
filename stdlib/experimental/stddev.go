package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	stddevSignature := runtime.MustLookupBuiltinType("experimental", "stddev")
	runtime.RegisterPackageValue("experimental", "stddev", flux.MustValue(flux.FunctionValue("stddev", universe.CreateStddevOpSpec, stddevSignature)))
}
