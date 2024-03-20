package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	histogramSignature := runtime.MustLookupBuiltinType("experimental", "histogram")
	runtime.RegisterPackageValue("experimental", "histogram", flux.MustValue(flux.FunctionValue("histogram", universe.CreateHistogramOpSpec, histogramSignature)))
}
