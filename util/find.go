package util

import (
	"fmt"
	"log"
)

// FindI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router, reporting back the location
func FindI2PIsInstalledDefaultLocation() (string, error) {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return I2PD_WINDOWS_DEFAULT_LOCATION, nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return I2PD_LINUX_SYSTEM_LOCATION[0], nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return I2PD_LINUX_SYSTEM_LOCATION[1], nil
	}
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return WINDOWS_DEFAULT_LOCATION, nil
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return WINDOWS_ZERO_NSIS_DEFAULT_LOCATION, nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[0], nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[1], nil
	}
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return HOME_DIRECTORY_LOCATION, nil
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return OSX_DEFAULT_LOCATION, nil
	}
	return "", fmt.Errorf("i2p router not found.")
}
