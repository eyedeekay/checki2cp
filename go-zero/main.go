package main

import (
	"log"
)

import "github.com/eyedeekay/checki2cp/zerobundle"

func main() {
	if err := zerobundle.UnpackZero(); err != nil {
		log.Println(err)
	}
	latest := zerobundle.LatestZero()
	log.Println("latest zero version is:", latest)
	if err := WriteI2CPConf(); err != nil {
		if ok, err := checki2p.ConditionallyLaunchI2P(); ok {
			if err != nil {
				log.Println(err)
			} else {
				if err := zerobundle.StartZero(); err != nil {
					log.Fatal(err)
				}
				log.Println("Starting SAM")
				if err := zerobundle.SAM(); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			log.Println("Undefined I2P launching error")
		}
	}
}
