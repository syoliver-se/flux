package iox

import (
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/internal/function"
	"github.com/syoliver-se/flux/values"
)

const pkgpath = "experimental/iox"

func init() {
	b := function.ForPackage(pkgpath)
	b.Register("from", func(args *function.Arguments) (values.Value, error) {
		return nil, errors.New(codes.Unimplemented, "iox.from() is not implemented outside cloud 2.x")
	})
	b.RegisterSource("sql", SqlKind, createSqlProcedureSpec)
}
