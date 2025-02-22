package checki2p

import (
	"fmt"
	"log"
	"os"
	"os/user"

	checksam "github.com/eyedeekay/checki2cp/samcheck"
	"github.com/eyedeekay/checki2cp/util"
	go_i2cp "github.com/go-i2p/go-i2cp"
)

func i2pdArgs() ([]string, error) {
	return []string{""}, nil
}

// CheckI2PIsRunning is frequently the only thing I need a reliable, non-SAM
// way to test at runtime.
func CheckI2PIsRunning() (bool, error) {
	log.Println("Trying to discover a running I2P router")
	if checksam.CheckSAMAvailable("127.0.0.1:7656") {
		return true, nil
	}
	client := go_i2cp.NewClient(nil)
	err := client.Connect()
	if err != nil {
		return false, nil
	}
	destination, err := go_i2cp.NewDestination()
	if err != nil {
		return false, err
	}
	if destination == nil {
		return false, fmt.Errorf("Client connection was made bug no destination could be generated.")
	}
	client.Disconnect()
	log.Println("I2P is running.")
	return true, nil
}

// CheckI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router.
func CheckI2PIsInstalledDefaultLocation() (bool, error) {
	if util.CheckFileExists(util.I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return true, nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return true, nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return true, nil
	}
	if util.CheckFileExists(util.WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return true, nil
	}
	if util.CheckFileExists(util.WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p zero router detected")
		return true, nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if util.CheckFileExists(util.HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if util.CheckFileExists(util.OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return true, nil
	}
	return false, nil
}

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
	if util.CheckFileExists(util.I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return "i2pd", nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if util.CheckFileExists(util.I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if util.CheckFileExists(util.WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2psvc", nil
	}
	if util.CheckFileExists(util.WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2p-zero", nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if util.CheckFileExists(util.LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if util.CheckFileExists(util.HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return UserFind(), nil
	}
	if util.CheckFileExists(util.OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return UserFind(), nil
	}
	return "", nil
}
