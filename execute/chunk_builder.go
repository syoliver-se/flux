package execute

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/array"
	"github.com/syoliver-se/flux/arrow"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/execute/table"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/memory"
	"github.com/syoliver-se/flux/values"
)

type ChunkBuilder struct {
	cols     []flux.ColMeta
	builders []array.Builder
}

func NewChunkBuilder(cols []flux.ColMeta, size int, mem memory.Allocator) *ChunkBuilder {
	builders := make([]array.Builder, len(cols))
	for i, col := range cols {
		b := arrow.NewBuilder(col.Type, mem)
		b.Resize(size)
		builders[i] = b
	}
	return &ChunkBuilder{cols: cols, builders: builders}
}

func (b *ChunkBuilder) AppendRecord(record values.Object) error {
	for i, col := range b.cols {
		v, ok := record.Get(col.Label)
		if !ok {
			return errors.Newf(codes.Internal, "could not find column %s in record", col.Label)
		}
		if err := arrow.AppendValue(b.builders[i], v); err != nil {
			return err
		}
	}
	return nil
}

func (b *ChunkBuilder) Build(key flux.GroupKey) table.Chunk {
	buf := arrow.TableBuffer{
		GroupKey: key,
		Columns:  b.cols,
	}
	vals := make([]array.Array, 0, len(b.builders))
	for _, builder := range b.builders {
		vals = append(vals, builder.NewArray())
	}
	buf.Values = vals
	return table.ChunkFromBuffer(buf)
}
