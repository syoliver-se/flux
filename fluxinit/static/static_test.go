package static_test

import (
	"testing"

	_ "github.com/syoliver-se/flux/fluxinit/static"
)

func TestBuiltins(t *testing.T) {
	t.Log("Testing that importing builtins does not panic")
}
