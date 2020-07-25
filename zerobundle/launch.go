package zerobundle

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

var command exec.Cmd

func GetZeroPID() int {
	return command.Process.Pid
}

func LatestZeroBinDir() string {
	var dir string
	var err error
	if dir, err = UnpackZeroDir(); err == nil {
		fs, er := ioutil.ReadDir(dir)
		if er != nil {
			log.Fatal(er)
		}
		return filepath.Join(dir, fs[len(fs)-1].Name(), "router", "bin")
	} else {
		log.Fatal(err)
	}
	return ""
}

func LatestZeroBinDirJavaHome() string {
	fs, er := ioutil.ReadDir(JAVA_I2P_OPT_DIR)
	if er != nil {
		log.Fatal(er)
	}
	return filepath.Join(JAVA_I2P_OPT_DIR, fs[len(fs)-1].Name(), "router", "bin")
}
