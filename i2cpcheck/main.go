package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eyedeekay/checki2cp"
	"github.com/eyedeekay/checki2cp/proxycheck"
	"github.com/eyedeekay/checki2cp/samcheck"
)

// CheckI2PRunning returns true if I2P is running. That's all.
func CheckI2PRunning(needI2CP, needProxy, needSAM bool) (bool, error) {
	if needI2CP {
		if notRunning, inError := checki2p.CheckI2PIsRunning(); inError != nil {
			return false, fmt.Errorf("A strange error occurred: %s", inError)
		} else if notRunning {
			return true, fmt.Errorf("I2P is already running with an open I2CP port")
		}
	}
	if needProxy {
		if checkproxy.ProxyDotI2P() {
			return true, fmt.Errorf("I2P is already running with an open HTTP proxy")
		}
	}
	if needSAM {
		if checksam.CheckSAMAvailable("") {
			return true, fmt.Errorf("I2P is already running with an open SAM API")
		}
	}
	return false, fmt.Errorf("I2P is not running.")
}

func main() {
	ok, err := CheckI2PRunning(true, true, true)
	if err != nil {
		if ok {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
	firewallport, err := checki2p.GetFirewallPort()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("I2P's firewall port is:", firewallport)

	ok, err = checki2p.CheckI2PIsInstalledDefaultLocation()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("I2P is installed, successfully confirmed")
		path, err := checki2p.FindI2PIsInstalledDefaultLocation()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("I2P is installed at:", path)
		os.Exit(0)
	} else {
		log.Println("I2P is not a default location, user feedback is needed")
		os.Exit(1)
	}

}
