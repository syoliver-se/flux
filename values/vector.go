package values

import (
	arrow "github.com/syoliver-se/flux/array"
	"github.com/syoliver-se/flux/semantic"
)

type Vector interface {
	Value
	ElementType() semantic.MonoType
	Arr() arrow.Array
	IsRepeat() bool
}
