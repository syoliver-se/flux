package experimental

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func init() {
	kamaSignature := runtime.MustLookupBuiltinType("experimental", "kaufmansAMA")
	runtime.RegisterPackageValue("experimental", "kaufmansAMA", flux.MustValue(flux.FunctionValue("kaufmansAMA", universe.CreatekamaOpSpec, kamaSignature)))
}
