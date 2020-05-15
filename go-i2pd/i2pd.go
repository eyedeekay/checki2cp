package main

import (
	"log"

	"github.com/eyedeekay/checki2cp/i2pdbundle"
)

func main() {
	if err := i2pd.UnpackI2Pd(); err != nil {
		log.Println(err)
	}
	if path, err := i2pd.FindI2Pd(); err != nil {
		log.Println(err)
	} else {
		log.Println(path)
	}
	//	if cmd, err := i2pd.LaunchI2Pd(); err != nil {
	if _, err := i2pd.LaunchI2Pd(); err != nil {
		log.Println(err)
	}
}

