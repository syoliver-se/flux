package values

import "github.com/syoliver-se/flux/semantic"

var (
	trueValue Value = value{
		t: semantic.BasicBool,
		v: true,
	}
	falseValue Value = value{
		t: semantic.BasicBool,
		v: false,
	}
)

func NewBool(v bool) Value {
	if v {
		return trueValue
	}
	return falseValue
}
