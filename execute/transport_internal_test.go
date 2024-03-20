package execute

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/execute/table"
)

func NewProcessMsg(tbl flux.Table) ProcessMsg {
	return &processMsg{table: tbl}
}

func NewProcessChunkMsg(chunk table.Chunk) ProcessChunkMsg {
	return &processChunkMsg{chunk: chunk}
}

func NewFlushKeyMsg(key flux.GroupKey) FlushKeyMsg {
	return &flushKeyMsg{key: key}
}

func NewFinishMsg(err error) FinishMsg {
	return &finishMsg{err: err}
}
