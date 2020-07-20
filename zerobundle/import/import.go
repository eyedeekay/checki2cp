package zerobundle

import (
	"io/ioutil"
	"os"
	"runtime"
	"path/filepath"

	"github.com/mholt/archiver/v3"

	"github.com/eyedeekay/checki2cp/zerobundle/parts/aa"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ab"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ac"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ad"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ae"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/af"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ag"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ah"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ai"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/aj"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/ak"
	"github.com/eyedeekay/checki2cp/zerobundle/parts/al"
)

func TBZBytes() ([]byte, error) {
	var bytes []byte
	ba, err := izaa.WriteBrowser(izaa.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, ba...)
	bb, err := izab.WriteBrowser(izab.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bb...)
	bc, err := izac.WriteBrowser(izac.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bc...)
	bd, err := izad.WriteBrowser(izad.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bd...)
	be, err := izae.WriteBrowser(izae.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, be...)
	bf, err := izaf.WriteBrowser(izaf.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bf...)
	bg, err := izag.WriteBrowser(izag.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bg...)
	bh, err := izah.WriteBrowser(izah.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bh...)
	bi, err := izai.WriteBrowser(izai.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bi...)
	bj, err := izaj.WriteBrowser(izaj.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bj...)
	bk, err := izak.WriteBrowser(izak.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bk...)
	bl, err := izal.WriteBrowser(izal.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bl...)
	return bytes, nil
}

func Write() error {
	var platform = "linux"
	if runtime.GOOS == "windows" {
		platform = "win"
	}
	if runtime.GOOS == "darwin" {
		platform = "mac"
	}
	bytes, err := TBZBytes()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("i2p-zero-"+platform+".v1.17.zip", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Unpack(destinationDirectory string) error {
	var platform = "linux"
	if runtime.GOOS == "windows" {
		platform = "win"
	}
	if runtime.GOOS == "darwin" {
		platform = "mac"
	}
	if destinationDirectory == "" {
		destinationDirectory = "."
	}
	err := os.RemoveAll(filepath.Join(destinationDirectory, "i2p-zero"))
	if err != nil {
		return err
	}
	err = Write()
	if err != nil {
		return err
	}
	err = archiver.Unarchive("i2p-zero-"+platform+".v1.17.zip", destinationDirectory)
	if err != nil {
		return err
	}
	return nil
}
