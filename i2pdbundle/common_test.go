package i2pd

import (
	"testing"
)

func TestFSListDirs(t *testing.T) {
    list, err := FindAllSubdirectories(FS)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(list)
}