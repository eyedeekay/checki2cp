package util

import (
	"log"
	"os"
	"os/user"
)

// UserFind makes sure that we never mis-identify the user account because of
// sudo
func UserFind() string {
	if os.Geteuid() == 0 {
		str := os.Getenv("SUDO_USER")
		return str
	}
	if un, err := user.Current(); err == nil {
		return un.Name
	}
	return ""
}

// CheckI2PUserName looks in various locations for the
// presence of an I2P router and guesses the username it
// should run under. On Windows it returns the EXE name.
func CheckI2PUserName() (string, error) {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return "i2pd", nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2psvc", nil
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2p-zero", nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return UserFind(), nil
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return UserFind(), nil
	}
	return "", nil
}
