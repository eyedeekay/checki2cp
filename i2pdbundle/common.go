package i2pd

import (
	"io/ioutil"
	"log"
	"os"
)

type fsi interface {
	IsDir() bool
	Readdir(int) ([]os.FileInfo, error)
	Open(name string) (http.File, error)
}

func FindAllDirectories(filesystem fsi) ([]string, error) {
	if filesystem.IsDir() {
		filelist, err := filesystem.Readdir(0)
		if err != nil {
			return nil, err
		}
		var rlist []string
		for index, file := range filelist {
			if file.IsDir() {
				rlist = append(rlist, file.Name())
				log.Println(index, file.Name())
			}
		}
	}
	return nil, nil
}

func FindAllFiles(filesystem fsi) ([]string, error) {
	if filesystem.IsDir() {
		filelist, err := filesystem.Readdir(0)
		if err != nil {
			return nil, err
		}
		var rlist []string
		for index, file := range filelist {
			if !file.IsDir() {
				rlist = append(rlist, file.Name())
				log.Println(index, file.Name())
			}
		}
	}
	return nil, nil
}

func WriteAllFiles(filesystem fsi, unpackdir string) error {
	if filesystem.IsDir() {
		if filelist, err := filesystem.Readdir(0); err == nil {
			var rlist []string
			for index, fi := range filelist {
				if file, err := filesystem.Open(fi.Name()); err == nil {
					if !fi.IsDir() {
						var buf []byte
						if _, err := file.Read(buf); err == nil {
							rlist = append(rlist, fi.Name())
							log.Println(index, fi.Name())
							if err := ioutil.WriteFile(unpackdir+"/"+fi.Name(), buf, fi.Mode()); err != nil {
								return err
							}
						} else {
							return err
						}
					}
					file.Close()
				} else {
					return err
				}
			}
		} else {
			return err
		}
	}
	return nil
}
