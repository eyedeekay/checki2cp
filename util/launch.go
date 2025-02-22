package util

import (
	"log"
)

// CheckI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router.
func CheckI2PIsInstalledDefaultLocation() (bool, error) {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return true, nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return true, nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return true, nil
	}
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return true, nil
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p zero router detected")
		return true, nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return true, nil
	}
	return false, nil
}
