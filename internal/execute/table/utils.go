package table

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/array"
	"github.com/syoliver-se/flux/execute/table"
)

func Values(cr flux.ColReader, j int) array.Array {
	return table.Values(cr, j)
}
