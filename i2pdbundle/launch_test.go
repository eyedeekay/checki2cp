package i2pd

import (
	"os/exec"
	"testing"
)

func TestUnpackI2PDPath(t *testing.T) {
	out, err := UnpackI2PdDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Testing output in:", out)
	err = WriteAllFiles(FS, out)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Listing output in parent", "ls", "-lahR", out)
	err = exec.Command("ls", "-lahR", out).Run()
	if err != nil {
		t.Fatal(err)
	}
	err = UnpackI2Pd()
	if err != nil {
		t.Fatal(err)
	}

}
