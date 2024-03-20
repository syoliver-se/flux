package execute

import (
	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/internal/execute/groupkey"
	"github.com/syoliver-se/flux/values"
)

func NewGroupKey(cols []flux.ColMeta, values []values.Value) flux.GroupKey {
	return groupkey.New(cols, values)
}
