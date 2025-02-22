package checki2p

import (
	"fmt"
	"log"

	"github.com/eyedeekay/checki2cp/util"
)

// FindI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router, reporting back the location
func FindI2PIsInstalledDefaultLocation() (string, error) {
	if util.CheckFileExists(util.I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return util.I2PD_WINDOWS_DEFAULT_LOCATION, nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return util.I2PD_LINUX_SYSTEM_LOCATION[0], nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return util.I2PD_LINUX_SYSTEM_LOCATION[1], nil
	}
	if util.CheckFileExists(util.WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return util.WINDOWS_DEFAULT_LOCATION, nil
	}
	if util.CheckFileExists(util.WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return util.WINDOWS_ZERO_NSIS_DEFAULT_LOCATION, nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return util.LINUX_SYSTEM_LOCATION[0], nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return util.LINUX_SYSTEM_LOCATION[1], nil
	}
	if util.CheckFileExists(util.HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return util.HOME_DIRECTORY_LOCATION, nil
	}
	if util.CheckFileExists(util.OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return util.OSX_DEFAULT_LOCATION, nil
	}
	return "", fmt.Errorf("i2p router not found.")
}
