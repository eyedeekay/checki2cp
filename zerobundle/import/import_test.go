package zerobundle

import (
	"testing"
)

func TestWriteTBZ(t *testing.T) {
	if err := Unpack(""); err != nil {
		t.Fatal(err)
	}
	t.Log("Success")
}
