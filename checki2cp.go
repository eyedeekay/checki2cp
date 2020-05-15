package checki2p

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/eyedeekay/go-i2cp"
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

func home() string {
	if runtime.GOOS == "windows" {
		return "\\i2p"
	}
	return "/.i2p"
}

var (
	I2CP_HOST                     string   = ""
	I2CP_PORT                     string   = ""
	WINDOWS_DEFAULT_LOCATION      string   = `C:\\Program Files\i2p\i2psvc.exe`
	I2PD_WINDOWS_DEFAULT_LOCATION string   = `C:\\Program Files\I2Pd\i2pd.exe`
	LINUX_SYSTEM_LOCATION         []string = []string{"/usr/bin/i2prouter", "/usr/sbin/i2prouter"}
	I2PD_LINUX_SYSTEM_LOCATION    string   = "/usr/sbin/i2pd"
	I2P_ASUSER_HOME_LOCATION      string   = inithome(home())
	HOME_DIRECTORY_LOCATION       string   = inithome("/i2p/i2prouter")
	OSX_DEFAULT_LOCATION          string   = inithome("/Library/Application Support/i2p/clients.config")
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
	if checkfileexists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return true, nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION[1]) {
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
	if checkfileexists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return "i2pd", nil
	}
	if checkfileexists(I2PD_LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2pd router detected")
		return "i2pd", nil
	}
	if checkfileexists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2psvc", nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return "i2psvc", nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION[1]) {
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

// GetFirewallPort finds the configured UDP port of your I2P router to help
// configure firewalls. It does this by finding the router.config and reading
// it. This function will not work on I2Pd routers yet but it should be easy
// to add support once I get some more time to test and research it.
func GetFirewallPort() (string, error) {
	log.Println(I2P_ASUSER_HOME_LOCATION)
	file, err := ioutil.ReadFile(I2P_ASUSER_HOME_LOCATION + "/router.config")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(file), "\n")
	for index, line := range lines {
		if strings.HasPrefix(line, "i2np.udp.port") {
			log.Println(line, index)
			return strings.Replace(line, "i2np.udp.port=", "", -1), nil
		}
	}
	return "", fmt.Errorf("Improperly formed router.config file")
}

// FindI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router, reporting back the location
func FindI2PIsInstalledDefaultLocation() (string, error) {
	if checkfileexists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return I2PD_WINDOWS_DEFAULT_LOCATION, nil
	}
	if checkfileexists(I2PD_LINUX_SYSTEM_LOCATION) {
		log.Println("Linux i2pd router detected")
		return I2PD_LINUX_SYSTEM_LOCATION, nil
	}
	if checkfileexists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return WINDOWS_DEFAULT_LOCATION, nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[0], nil
	}
	if checkfileexists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[1], nil
	}
	if checkfileexists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return HOME_DIRECTORY_LOCATION, nil
	}
	if checkfileexists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return OSX_DEFAULT_LOCATION, nil
	}
	return "", fmt.Errorf("i2p router not found.")
}

// ConditionallyLaunchI2P If an already-installed I2P router is present, then
// make sure that it is started, i.e. launch the router *only* if it is not
// already running.
func ConditionallyLaunchI2P() (bool, error) {
	ok, err := CheckI2PIsInstalledDefaultLocation()
	if err != nil {
		return false, err
	}
	if ok {
		if ok, err := CheckI2PIsRunning(); err == nil {
			if !ok {
				path, err := FindI2PIsInstalledDefaultLocation()
				if err != nil {
					return false, err
				}
				if strings.HasSuffix(path, "i2prouter") {
					cmd := exec.Command(path, "start")
					if err := cmd.Start(); err != nil {
						return false, fmt.Errorf("I2P router startup failure %s", err)
					}
				} else {
					cmd := exec.Command(path, "--daemon")
					if err := cmd.Start(); err != nil {
						return false, fmt.Errorf("I2P router startup failure %s", err)
					}
				}
				return true, nil
			}
			return true, nil
		} else {
			return false, err
		}
	} else {
		return false, fmt.Errorf("I2P is not a default location, please set $I2P environment variable")
	}
}
