package zerobundle

import (
	//	"io/ioutil"
	//	"log"

	"path/filepath"
)

func LatestZero() string {
	return filepath.Join(LatestZeroBinDir(), "i2p-zero")
}

func LatestZeroJavaHome() string {
	return filepath.Join(LatestZeroBinDirJavaHome(), "i2p-zero")
}

func RunZero() {

}

func RunZeroJavaHome() {

}

func StopZero() {

}
