package arrow

import (
	arrowmemory "github.com/apache/arrow/go/v7/arrow/memory"
	"github.com/syoliver-se/flux/memory"
)

func NewAllocator(a memory.Allocator) arrowmemory.Allocator {
	return a
}
