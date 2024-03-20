package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	quantileSignature := runtime.MustLookupBuiltinType("experimental", "quantile")
	runtime.RegisterPackageValue("experimental", "quantile", flux.MustValue(flux.FunctionValue("quantile", universe.CreateQuantileOpSpec, quantileSignature)))
}
