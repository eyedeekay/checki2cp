package zerobundle

import (
	//	"bytes"
	//	"fmt"
	//	"io"
	"log"
	"os"
	"path/filepath"

	. "github.com/eyedeekay/checki2cp/zerobundle/import"
)

var I2P_DIRECTORY_PATH = os.Getenv("I2P_DIRECTORY_PATH")

func userFind() string {
	if os.Geteuid() == 0 {
		log.Fatal("Do not run this application as root!")
	}
	if un, err := os.UserHomeDir(); err == nil {
		os.MkdirAll(filepath.Join(un, "i2p"), 0755)
		return un
	}
	return ""
}

var JAVA_I2P_OPT_DIR = filepath.Join(userFind(), "/i2p/opt/i2p-zero")

// UnpackI2PdDir tells the bundle where the I2Pd-bundling app exists
func UnpackZeroDir() (string, error) {
	if I2P_DIRECTORY_PATH != "" {
		return I2P_DIRECTORY_PATH, nil
	}
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	ri2pd := filepath.Dir(executablePath)
	return ri2pd, nil
}

func UnpackZero() error {
	var dir string
	var err error
	if dir, err = UnpackZeroDir(); err == nil {
		if err := Unpack(dir); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func UnpackZeroJavaHome() error {
	if err := Unpack(JAVA_I2P_OPT_DIR); err != nil {
		return err
	}
	return nil
}
