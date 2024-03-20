package universe

import (
	"context"

	"github.com/apache/arrow/go/v7/arrow/memory"
	"github.com/syoliver-se/flux/execute"
)

func NewMapTransformation(ctx context.Context, id execute.DatasetID, spec *MapProcedureSpec, mem memory.Allocator) (execute.Transformation, execute.Dataset, error) {
	return newMapTransformation(ctx, id, spec, mem)
}
