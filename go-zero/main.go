package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

import (
	"github.com/eyedeekay/checki2cp"
	"github.com/eyedeekay/checki2cp/zerobundle"
)

var i2cpConf = `i2cp.tcp.host=127.0.0.1
i2cp.tcp.port=7654
`

func WriteI2CPConf() error {
	dir, err := zerobundle.UnpackZeroDir()
	if err != nil {
		return err
	}
	os.Setenv("I2CP_HOME", dir)
	os.Setenv("GO_I2CP_CONF", "/.i2cp.conf")
	home := os.Getenv("I2CP_HOME")
	conf := os.Getenv("GO_I2CP_CONF")
	if err := ioutil.WriteFile(filepath.Join(home, conf), []byte(i2cpConf), 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := WriteI2CPConf(); err != nil {
    log.Println(err)
	}
	if ok, err := checki2p.ConditionallyLaunchI2P(); ok {
		if err != nil {
			log.Println(err)
		}
	} else {
		if err := zerobundle.UnpackZero(); err != nil {
			log.Println(err)
		}
		latest := zerobundle.LatestZero()
		log.Println("latest zero version is:", latest)
		if err := zerobundle.StartZero(); err != nil {
			log.Fatal(err)
		}
		log.Println("Starting SAM")
		if err := zerobundle.SAM(); err != nil {
			log.Fatal(err)
		}
		log.Println("Undefined I2P launching error")
	}
}
