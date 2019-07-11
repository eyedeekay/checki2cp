package main

import (
	"github.com/eyedeekay/checki2cp"
	"log"
)

func main() {
	ok, err := checki2p.CheckI2PIsRunning()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("I2P is running, successfully confirmed I2CP")
		return
	} else {
		log.Println("I2P is not running, further testing is needed")
		return
	}
	ok, err = checki2p.CheckI2PIsInstalledDefaultLocation()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("I2P is installed, successfully confirmed")
		return
	} else {
		log.Println("I2P is in a default location, user feedback is needed")
		return
	}
}
