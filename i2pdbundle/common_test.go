package i2pd

import (
	"testing"
)

func TestFSListFiles(t *testing.T) {
    list, err := FindAllFiles(FS)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(list)
}

func TestFSListDirs(t *testing.T) {
    list, err := FindAllDirectories(FS)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(list)
}
