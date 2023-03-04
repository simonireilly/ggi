package files

import (
	"testing"
)

func TestList(t *testing.T) {
	a := NewAdapter()
	res := a.List()

	if len(res) < 1 {
		t.Fatalf("expected greater than one files in gitignore embedded system")
	}
}
