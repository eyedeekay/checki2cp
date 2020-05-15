package main

import (
	"flag"
	"log"

	"github.com/eyedeekay/checki2cp/i2pdbundle"
)

func main() {
	boolPtr := flag.Bool("force", false, "Force an embedded I2Pd router to start")
	flag.Parse()
	if err := i2pd.UnpackI2Pd(); err != nil {
		log.Println(err)
	}
	if path, err := i2pd.FindI2Pd(); err != nil {
		log.Println(err)
	} else {
		log.Println(path)
	}
	if !*boolPtr {
		//	if cmd, err := i2pd.LaunchI2Pd(); err != nil {
		if _, err := i2pd.LaunchI2Pd(); err != nil {
			log.Println(err)
		}
	} else {
		if _, err := i2pd.LaunchI2PdForce(); err != nil {
			log.Println(err)
		}
	}
}
