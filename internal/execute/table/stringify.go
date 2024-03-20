package table

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/execute/table"
)

// Stringify will read a table and turn it into a human-readable string.
func Stringify(t flux.Table) string {
	return table.Stringify(t)
}
