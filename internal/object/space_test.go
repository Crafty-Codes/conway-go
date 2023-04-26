package object

import (
	"testing"
)

func TestNewSpace(t *testing.T) {
	spaceObj := NewSpace(4, 6)

	if len(spaceObj) != 4 || len(spaceObj[0]) != 6 {
		t.Error("object.NewSpace(4, 6) didin't produce a 4x6 split")
	}
}
