package boolean

import (
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/values"
)

func init() {
	runtime.RegisterPackageValue("internal/boolean", "true", values.NewBool(true))
	runtime.RegisterPackageValue("internal/boolean", "false", values.NewBool(false))
}
