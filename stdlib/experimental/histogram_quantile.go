package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	histogramQuantileSignature := runtime.MustLookupBuiltinType("experimental", "histogramQuantile")
	runtime.RegisterPackageValue("experimental", "histogramQuantile", flux.MustValue(flux.FunctionValue("histogramQuantile", universe.CreateHistogramQuantileOpSpec, histogramQuantileSignature)))
}
