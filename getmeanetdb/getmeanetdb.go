package getmeanetdb

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eyedeekay/checki2cp"
	"github.com/eyedeekay/checki2cp/i2pdbundle"
)

var UnpackI2PdDir = i2pd.UnpackI2PdDir

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

	}
	switch path {
	case checki2p.WINDOWS_DEFAULT_LOCATION:
		env := os.Getenv("LOCALAPPDATA")
		if env == "" {
			return "", fmt.Errorf("Could not find local appdata path")
		}
		return filepath.Join(env, "I2P"), nil
	case checki2p.I2PD_WINDOWS_DEFAULT_LOCATION:
		env := os.Getenv("LOCALAPPDATA")
		if env == "" {
			return "", fmt.Errorf("Could not find local appdata path")
		}
		return filepath.Join(env, "i2pd"), nil
	case checki2p.LINUX_SYSTEM_LOCATION[0]:
		return "/var/lib/i2p/i2p-config/netDb", nil
	case checki2p.LINUX_SYSTEM_LOCATION[1]:
		return "/var/lib/i2p/i2p-config/netDb", nil
	case checki2p.I2PD_LINUX_SYSTEM_LOCATION:
	case checki2p.I2P_ASUSER_HOME_LOCATION:
		return checki2p.I2P_ASUSER_HOME_LOCATION, nil
	case checki2p.HOME_DIRECTORY_LOCATION:
		return checki2p.I2P_ASUSER_HOME_LOCATION, nil
	case checki2p.OSX_DEFAULT_LOCATION:
		return checki2p.I2P_ASUSER_HOME_LOCATION, nil
	}
	return UnpackI2PdDir()
}
