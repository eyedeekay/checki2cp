package main

import (
	"github.com/eyedeekay/checki2cp"
	"log"
	"os"
)

func main() {
	ok, err := checki2p.CheckI2PIsRunning()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("I2P is running, successfully confirmed I2CP")
		os.Exit(0)
	} else {
		log.Println("I2P is not running, further testing is needed")
	}
	ok, err = checki2p.CheckI2PIsInstalledDefaultLocation()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("I2P is installed, successfully confirmed")
		os.Exit(0)
	} else {
		log.Println("I2P is not a default location, user feedback is needed")
		os.Exit(1)
	}
}
