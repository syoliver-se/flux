package runtime

import (
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/runtime"
	"github.com/syoliver-se/flux/values"
)

const versionFuncName = "version"

var errBuildInfoNotPresent = errors.New(codes.NotFound, "build info is not present")

func init() {
	runtime.RegisterPackageValue("runtime", versionFuncName, values.NewFunction(
		versionFuncName,
		runtime.MustLookupBuiltinType("runtime", versionFuncName),
		Version,
		false,
	))
}
