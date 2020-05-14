package i2pd

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/shurcooL/httpfs/vfsutil"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FileOK(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return err
	} else {
		return err
	}
}

//var walkFn = func(path string, fi os.FileInfo, err error) error {
var walkFn = func(path string, fi os.FileInfo, r io.ReadSeeker, err error) error {
	if err != nil {
		log.Printf("can't stat file %s: %v\n", path, err)
		return nil
	}
	fmt.Println(path)
	if !fi.IsDir() {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Printf("can't read file %s: %v\n", path, err)
			return err
		}
		dir, err := UnpackI2PdDir()
		if err != nil {
			log.Printf("can't find path: %v\n", err)
			return err
		}
		err = ioutil.WriteFile(filepath.Join(dir, path), b, fi.Mode())
		if err != nil {
			log.Printf("can't write file %s: %v\n", filepath.Join(dir, path), err)
			return err
		}
		//fmt.Printf("%q\n", b)
		dirpath := strings.Split(path, ".")[0]
		log.Printf("wrote file %s: %v", filepath.Join(dir, path), fi.Mode())
		err = archiver.Unarchive(filepath.Join(dir, path), filepath.Join(dir, dirpath))
		if err != nil {
			log.Printf("can't unarchive file %s: %v\n", filepath.Join(dir, path), err)
			return err
		}
		log.Printf("unpacked file %s", filepath.Join(dir, path))
	}
	return nil
}

func WriteAllFiles(targetdir string) error {
	err := vfsutil.WalkFiles(FS, "/", walkFn)
	if err != nil {
		return err
	}
	return nil
}

func UnpackI2PdPath() (string, error) {
	dirPath, err := UnpackI2PdDir()
	if err != nil {
		return "", err
	}
	ri2pd := filepath.Join(dirPath, "i2pd")
	/*if err := FileOK(ri2pd); err != nil {
		return "", err
	}*/
	return ri2pd, nil
}

func UnpackI2PdLibPath() (string, error) {
	dirPath, err := UnpackI2PdDir()
	if err != nil {
		return "", err
	}
	rlib := filepath.Join(dirPath, "lib")
	/*if err := FileOK(rlib); err != nil {
		return "", err
	}*/
	return rlib, nil
}

func UnpackI2PdDir() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	ri2pd := filepath.Dir(executablePath)
	/*if err := FileOK(ri2pd); err != nil {
		return "", err
	}*/
	return ri2pd, nil
}
