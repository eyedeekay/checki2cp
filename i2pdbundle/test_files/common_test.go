package i2pdtest

import (
	"github.com/eyedeekay/checki2cp/i2pdbundle"
	"testing"
)

func TestFSListFiles(t *testing.T) {
	list, err := i2pd.FindAllFiles(FS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}

func TestFSListDirs(t *testing.T) {
	list, err := i2pd.FindAllDirectories(FS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}

func TestFSWriteFiles(t *testing.T) {
	err := i2pd.WriteAllFiles(FS, "test-test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Wrote files")
}
