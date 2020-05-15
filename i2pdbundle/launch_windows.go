package i2pd

import (
	"fmt"
	"log"
	//"os"
	"os/exec"
	"path/filepath"

	"github.com/eyedeekay/checki2cp"
	"github.com/eyedeekay/checki2cp/proxycheck"
	"github.com/eyedeekay/checki2cp/samcheck"
)

// FindI2P returns the absolute path to the i2pd executable
func FindI2Pd() (string, error) {
	path, err := UnpackI2PdPath()
	if err != nil {
		return "", err
	}
	if err := FileOK(filepath.Join(filepath.Dir(path), "i2pd.exe")); err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(path), "i2pd.exe"), nil
}

// UnpackI2Pd unpacks a working version of i2pd and some supporting libraries to a the directory of the executable
// that will start i2pd
func UnpackI2Pd() error {
	dir, err := UnpackI2PdPath()
	if err != nil {
		return err
	}
	err = WriteAllFiles(dir)
	if err != nil {
		return err
	}
	return nil
}

// LaunchI2PConditional will check for specific services and if they are not found, start a standalone router
func LaunchI2PdConditional(needI2CP, needProxy, needSAM bool) (*exec.Cmd, error) {
	if needI2CP {
		if notRunning, inError := checki2p.CheckI2PIsRunning(); inError != nil {
			return nil, inError
		} else if notRunning {
			return nil, fmt.Errorf("I2P is already running with an open I2CP port")
		}
	}
	if needProxy {
		if checkproxy.ProxyDotI2P() {
			return nil, fmt.Errorf("I2P is already running with an open HTTP proxy")
		}
	}
	if needSAM {
		if checksam.CheckSAMAvailable("") {
			return nil, fmt.Errorf("I2P is already running with an open SAM API")
		}
	}
	return LaunchI2PdForce()
}

// LaunchI2Pd will look for a running I2P router and if one is not found, it will start the embedded I2P router
func LaunchI2Pd() (*exec.Cmd, error) {
	return LaunchI2PdConditional(true, true, true)
}

// LaunchI2Pd attempts to launch the embedded I2P router no matter what.
func LaunchI2PdForce() (*exec.Cmd, error) {
	libPath, err := UnpackI2PdLibPath()
	if err != nil {
		return nil, err
	}
	if err := FileOK(libPath); err != nil {
		return nil, err
	}
	i2pd, err := FindI2Pd()
	if err := FileOK(libPath); err != nil {
		return nil, err
	}
	//err = os.Setenv("LD_LIBRARY_PATH", libPath)
	//if err != nil {
	//return nil, err
	//}
	log.Println(i2pd)
	runDir, err := UnpackI2PdDir()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(
		i2pd,
		"--datadir="+runDir,
		"--conf="+filepath.Join(runDir, "i2pd.conf"),
		"--tunconf="+filepath.Join(runDir, "tunnels.conf"),
		"--log=none",
	)
	//cmd.Env = append(os.Environ(),
	//"LD_LIBRARY_PATH="+libPath, // ignored
	//)
	log.Printf("running command: %v %s", cmd.Env, cmd.String())
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return cmd, nil
}
