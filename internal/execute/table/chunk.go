package table

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/arrow"
	"github.com/syoliver-se/flux/execute/table"
)

type Chunk = table.Chunk

func ChunkFromBuffer(buf arrow.TableBuffer) Chunk {
	return table.ChunkFromBuffer(buf)
}

func ChunkFromReader(cr flux.ColReader) Chunk {
	return table.ChunkFromReader(cr)
}
