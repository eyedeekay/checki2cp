package checksam

import "testing"

func TestCheckSAMAvailable(t *testing.T) {
	if CheckSAMAvailable("") {
		t.Log("SAM success")
	} else {
		t.Fatal("SAM not found")
	}
}
