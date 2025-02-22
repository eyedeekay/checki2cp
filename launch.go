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
	log.Println("Checking if I2P is installed at the default location.")
	ok, err := util.FindI2PIsInstalledDefaultLocation()
	if err != nil {
		return false, err
	}
	log.Println("I2P was found at the default location, proceeding with the launch procedure:", ok)
	if ok != "" {
		path := ok
		isRunning, err := CheckI2PIsRunning()
		if err != nil {
			return false, err
		}
		if !isRunning {
			return LaunchI2P(path)
		} else {
			log.Println("I2P is already running, no further action required.")
		}
		return true, nil
	}
	return false, fmt.Errorf("I2P is not found at the default location. Please set the $I2P environment variable")
}

// LaunchI2P starts an I2P router at the specified path.
// This function is used by ConditionallyLaunchI2P to start the router if it is not already running.
// it returns a boolean value indicating whether the router was started successfully or not.
// it also returns an error if the router could not be started.
func LaunchI2P(path string) (bool, error) {
	log.Println("Looking for an I2P router to start")
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
}
