package checki2p

import (
	"fmt"
	"github.com/eyedeekay/go-i2cp"
	"log"
	"os"
	"strings"
)

func inithome(str string) string {
	s, e := os.UserHomeDir()
	if e != nil {
		panic(e)
	}
	return s + str
}

func checkfileexists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

var (
	I2CP_HOST                     string = ""
	I2CP_PORT                     string = ""
	WINDOWS_DEFAULT_LOCATION      string = `C:\\Program Files\i2p\i2psvc.exe`
	I2PD_WINDOWS_DEFAULT_LOCATION string = `C:\\Program Files\I2Pd\i2pd.exe`
	LINUX_SYSTEM_LOCATION         string = "/usr/bin/i2prouter"
	I2PD_LINUX_SYSTEM_LOCATION    string = "/usr/sbin/i2pd"
	HOME_DIRECTORY_LOCATION       string = inithome("/i2p/i2prouter")
	OSX_DEFAULT_LOCATION          string = inithome("/Library/Application Support/i2p/clients.config")
)

// CheckIC2PIsRunning is frequently the only thing I need a reliable, non-SAM
// way to test at runtime.
func CheckI2PIsRunning() (bool, error) {
	client := go_i2cp.NewClient(nil)
	err := client.Connect()
	if err != nil {
		return false, err
	}
	destination, err := go_i2cp.NewDestination()
	if err != nil {
		return false, err
	}
	if destination == nil {
		return false, fmt.Errorf("")
	}
	client.Disconnect()
	return true, nil
}

// CheckI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router.
func CheckI2PIsInstalledDefaultLocation() (bool, error) {
	if checkfileexists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return true, nil
	}
	if checkfileexists(I2PD_LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2pd router detected")
		return true, nil
	}
	if checkfileexists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return true, nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if checkfileexists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if checkfileexists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return true, nil
	}
	return false, nil
}

func UserFind() string {
	if os.Geteuid() == 0 {
		str := os.Getenv("SUDO_USER")
		return str
	}
	if str, err := os.UserHomeDir(); err == nil {
		x := strings.Split(str, "/")
		return strings.Replace(x[len(x)-1], "/", "", -1)
	}
	return ""
}

// CheckI2PUserName looks in various locations for the
// presence of an I2P router and guesses the username it
// should run under. On Windows it returns the EXE name.
func CheckI2PUserName() (string, error) {
	if checkfileexists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return "i2pd.exe", nil
	}
	if checkfileexists(I2PD_LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if checkfileexists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2psvc.exe", nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if checkfileexists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return UserFind(), nil
	}
	if checkfileexists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return UserFind(), nil
	}
	return "", nil
}
