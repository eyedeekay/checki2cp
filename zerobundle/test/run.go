//+build run

package main

import (
	. "github.com/eyedeekay/checki2cp/zerobundle"
	"log"
)

func main() {
	if err := UnpackZeroJavaHome(); err != nil {
		log.Println(err)
	}
	latest := LatestZeroJavaHome()
	log.Println("latest zero version is:", latest)
	if err := RunZeroJavaHome(); err != nil {
		log.Fatal(err)
	}
}
