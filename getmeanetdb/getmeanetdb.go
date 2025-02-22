package getmeanetdb

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	checki2p "github.com/eyedeekay/checki2cp"
	"github.com/eyedeekay/checki2cp/util"
)

// WhereIsTheNetDB returns the path to whatever the first local NetDB
// it can find is. If it can't find one, it will output the path to an embedded NetDB
func WhereIstheNetDB() (string, error) {
	path, err := WhereIsTheConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "netDb"), nil
}

// WhereIstheConfigDir returns the path to the I2P configuration
// directory or the path to the one that will be created if  an
// embedded router is run.
func WhereIsTheConfigDir() (string, error) {
	path, err := checki2p.FindI2PIsInstalledDefaultLocation()
	if err != nil {
		log.Println("Warning: I2P is not installed at a default location.")
	}
	switch path {
	case util.WINDOWS_DEFAULT_LOCATION:
		env := os.Getenv("APPDATA")
		if env == "" {
			return "", fmt.Errorf("Could not find local appdata path")
		} else {
			env = os.Getenv("LOCALAPPDATA")
			if env == "" {
				return "", fmt.Errorf("Could not find local appdata path")
			}
		}
		return filepath.Join(env, "I2P"), nil
	case util.I2PD_WINDOWS_DEFAULT_LOCATION:
		env := os.Getenv("APPDATA")
		if env == "" {
			return "", fmt.Errorf("Could not find local appdata path")
		} else {
			env = os.Getenv("LOCALAPPDATA")
			if env == "" {
				return "", fmt.Errorf("Could not find local appdata path")
			}
		}
		return filepath.Join(env, "i2pd"), nil
	case util.LINUX_SYSTEM_LOCATION[0]:
		return "/var/lib/i2p/i2p-config/", nil
	case util.LINUX_SYSTEM_LOCATION[1]:
		return "/var/lib/i2p/i2p-config/", nil
	case util.I2PD_LINUX_SYSTEM_LOCATION[0]:
		return "/var/lib/i2pd/", nil
	case util.I2PD_LINUX_SYSTEM_LOCATION[1]:
		return "/var/lib/i2pd/", nil
	case util.I2P_ASUSER_HOME_LOCATION:
		return util.I2P_ASUSER_HOME_LOCATION, nil
	case util.HOME_DIRECTORY_LOCATION:
		return util.I2P_ASUSER_HOME_LOCATION, nil
	case util.OSX_DEFAULT_LOCATION:
		return util.I2P_ASUSER_HOME_LOCATION, nil
	}
	return "", fmt.Errorf("Could not find local I2P configuration directory")
}
