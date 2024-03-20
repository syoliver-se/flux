package sql

import (
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
)

// ErrorDriverDisabled indicates a given database driver is disabled.
var ErrorDriverDisabled = errors.New(codes.Unimplemented, "database driver disabled")
