// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: constant.gen.go.tmpl

package arrowutil

import (
	"fmt"

	"github.com/syoliver-se/flux/array"
)

func IsConstant(arr array.Array) bool {
	switch arr := arr.(type) {
	case *array.Int:
		return IsIntConstant(arr)
	case *array.Uint:
		return IsUintConstant(arr)
	case *array.Float:
		return IsFloatConstant(arr)
	case *array.Boolean:
		return IsBooleanConstant(arr)
	case *array.String:
		return IsStringConstant(arr)

	default:
		panic(fmt.Errorf("unsupported array datat ype: %s", arr.DataType()))
	}
}

func IsIntConstant(arr *array.Int) bool {
	// If all values are null, then that is still constant.
	if arr.NullN() == arr.Len() {
		return true
	} else if arr.NullN() > 0 {
		// At least one value is null, but not all so
		// not constant by definition.
		return false
	}

	// All values are non-null so check if they are all the same.
	v := arr.Value(0)
	for i, n := 1, arr.Len(); i < n; i++ {
		if arr.Value(i) != v {
			return false
		}
	}
	return true

}

func IsUintConstant(arr *array.Uint) bool {
	// If all values are null, then that is still constant.
	if arr.NullN() == arr.Len() {
		return true
	} else if arr.NullN() > 0 {
		// At least one value is null, but not all so
		// not constant by definition.
		return false
	}

	// All values are non-null so check if they are all the same.
	v := arr.Value(0)
	for i, n := 1, arr.Len(); i < n; i++ {
		if arr.Value(i) != v {
			return false
		}
	}
	return true

}

func IsFloatConstant(arr *array.Float) bool {
	// If all values are null, then that is still constant.
	if arr.NullN() == arr.Len() {
		return true
	} else if arr.NullN() > 0 {
		// At least one value is null, but not all so
		// not constant by definition.
		return false
	}

	// All values are non-null so check if they are all the same.
	v := arr.Value(0)
	for i, n := 1, arr.Len(); i < n; i++ {
		if arr.Value(i) != v {
			return false
		}
	}
	return true

}

func IsBooleanConstant(arr *array.Boolean) bool {
	// If all values are null, then that is still constant.
	if arr.NullN() == arr.Len() {
		return true
	} else if arr.NullN() > 0 {
		// At least one value is null, but not all so
		// not constant by definition.
		return false
	}

	// All values are non-null so check if they are all the same.
	v := arr.Value(0)
	for i, n := 1, arr.Len(); i < n; i++ {
		if arr.Value(i) != v {
			return false
		}
	}
	return true

}

func IsStringConstant(arr *array.String) bool {
	// If all values are null, then that is still constant.
	if arr.NullN() == arr.Len() {
		return true
	} else if arr.NullN() > 0 {
		// At least one value is null, but not all so
		// not constant by definition.
		return false
	}

	return arr.IsConstant()

}
