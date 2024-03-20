package execute

import (
	"io"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/values"
)

type RowReader interface {
	Next() bool
	GetNextRow() ([]values.Value, error)
	ColumnNames() []string
	ColumnTypes() []flux.ColType
	SetColumns([]interface{})
	io.Closer
}
