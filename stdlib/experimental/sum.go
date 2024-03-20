package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	sumSignature := runtime.MustLookupBuiltinType("experimental", "sum")
	runtime.RegisterPackageValue("experimental", "sum", flux.MustValue(flux.FunctionValue("sum", universe.CreateSumOpSpec, sumSignature)))
}
