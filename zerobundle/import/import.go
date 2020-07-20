package zerobundle

import (
	"io/ioutil"
	"os"
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
	ba, err := gsaa.WriteBrowser(gsaa.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, ba...)
	bb, err := gsab.WriteBrowser(gsab.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bb...)
	bc, err := gsac.WriteBrowser(gsac.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bc...)
	bd, err := gsad.WriteBrowser(gsad.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bd...)
	be, err := gsae.WriteBrowser(gsae.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, be...)
	bf, err := gsaf.WriteBrowser(gsaf.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bf...)
	bg, err := gsag.WriteBrowser(gsag.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bg...)
	bh, err := gsah.WriteBrowser(gsah.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bh...)
	bi, err := gsai.WriteBrowser(gsai.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bi...)
	bj, err := gsaj.WriteBrowser(gsaj.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bj...)
	bk, err := gsak.WriteBrowser(gsak.FS)
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, bk...)
	bl, err := gsal.WriteBrowser(gsal.FS)
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
	err = WriteTBZ()
	if err != nil {
		return err
	}
	err = archiver.Unarchive("i2p-zero-"+platform+".v1.17.zip", destinationDirectory)
	if err != nil {
		return err
	}
	return nil
}
