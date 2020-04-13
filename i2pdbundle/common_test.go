package i2pd

import (
	"testing"
    "github.com/eyedeekay/checki2cp/i2pdbundle/test_files"
)

func TestFSListFiles(t *testing.T) {
    list, err := FindAllFiles(i2pdtest.FS)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(list)
}

func TestFSListDirs(t *testing.T) {
    list, err := FindAllDirectories(i2pdtest.FS)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(list)
}
