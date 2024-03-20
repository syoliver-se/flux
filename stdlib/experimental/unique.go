package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	uniqueSignature := runtime.MustLookupBuiltinType("experimental", "unique")
	runtime.RegisterPackageValue("experimental", "unique", flux.MustValue(flux.FunctionValue("unique", universe.CreateUniqueOpSpec, uniqueSignature)))
}
