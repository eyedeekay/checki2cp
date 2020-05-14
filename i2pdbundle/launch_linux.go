package i2pd

import (
	"log"
	"os"
	"path/filepath"
	//"github.com/mholt/archiver/v3"
)

func FindI2Pd() (string, error) {
	path, err := UnpackI2PdPath()
	if err != nil {
		return "", err
	}
	if err := FileOK(filepath.Join(path, "i2pd")); err != nil {
		return "", err
	}
	return filepath.Join(path, "i2pd"), nil
}

func UnpackI2Pd() error {
	dir, err := UnpackI2PdPath()
	if err != nil {
		return err
	}
	err = WriteAllFiles(dir)
	if err != nil {
		return err
	}
	/*	log.Println("arc", dir+"/i2pd.tar.gz", dir+"/i2pd.tar")
		err = archiver.Unarchive(dir+"/i2pd.tar.gz", dir+"/i2pd.tar")
		if err != nil {
			return err
		}
		log.Println("arc", dir+"/i2pd.tar", dir+"/i2pd")
		err = archiver.Unarchive(dir+"/i2pd.tar", dir+"/i2pd")
		if err != nil {
			return err
		}*/
	return nil
}

func LaunchI2Pd() error {
	libPath, err := UnpackI2PdLibPath()
	if err != nil {
		return err
	}
	if err := FileOK(libPath); err != nil {
		return err
	}
	i2pd, err := FindI2Pd()
	if err := FileOK(libPath); err != nil {
		return err
	}
	err = os.Setenv("LD_LIBRARY_PATH", libPath)
	if err != nil {
		return err
	}
	log.Println(i2pd)
	return nil
}
