package checki2p

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/eyedeekay/checki2cp/util"
)

// ConditionallyLaunchI2P If an already-installed I2P router is present, then
// make sure that it is started, i.e. launch the router *only* if it is not
// already running.
func ConditionallyLaunchI2P() (bool, error) {
	log.Println("Checking if I2P is installed at a default location.")
	ok, err := util.FindI2PIsInstalledDefaultLocation()
	if err != nil {
		return false, err
	}
	log.Println("I2P was found at a default location, continuing procedure on:", ok)
	if ok != "" {
		ok, err := CheckI2PIsRunning()
		if err == nil {
			if !ok {
				log.Println("Looking for an I2P router to start")
				path, err := util.FindI2PIsInstalledDefaultLocation()
				if err != nil {
					return false, err
				}
				if strings.HasSuffix(path, "i2prouter") || strings.HasSuffix(path, "i2prouter.exe") || strings.HasSuffix(path, "i2psvc") || strings.HasSuffix(path, "i2psvc.exe") {
					cmd := exec.Command(path, "start")
					if err := cmd.Start(); err != nil {
						return false, fmt.Errorf("I2P router startup failure %s", err)
					}
				} else if strings.HasSuffix(path, "i2pd") || strings.HasSuffix(path, "i2pd.exe") {
					cmd := exec.Command(path, "--daemon")
					if err := cmd.Start(); err != nil {
						return false, fmt.Errorf("i2pd router startup failure %s", err)
					}
				} else {
					cmd := exec.Command(path)
					if err := cmd.Start(); err != nil {
						return false, fmt.Errorf("I2P Zero router startup failure %s", err)
					}
				}
				return true, nil
			} else {
				log.Println("I2P appears to be running, nothing to do.")
			}
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("I2P is not a default location, please set $I2P environment variable")
}
