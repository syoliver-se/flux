package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	skewSignature := runtime.MustLookupBuiltinType("experimental", "skew")
	runtime.RegisterPackageValue("experimental", "skew", flux.MustValue(flux.FunctionValue("skew", universe.CreateSkewOpSpec, skewSignature)))
}
