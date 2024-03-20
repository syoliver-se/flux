package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	fillSignature := runtime.MustLookupBuiltinType("experimental", "fill")
	runtime.RegisterPackageValue("experimental", "fill", flux.MustValue(flux.FunctionValue("fill", universe.CreateFillOpSpec, fillSignature)))
}
