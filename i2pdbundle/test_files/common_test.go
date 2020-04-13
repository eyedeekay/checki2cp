package i2pd

import (
	"testing"
    "github.com/eyedeekay/checki2cp/i2pdbundle"
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
