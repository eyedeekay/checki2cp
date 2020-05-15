package i2pd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v3"
	"github.com/shurcooL/httpfs/vfsutil"
)

var configFile = `## Configuration file for a typical i2pd user
## See https://i2pd.readthedocs.org/en/latest/configuration.html
## for more options you can use in this file.

#log = file
#logfile = ./i2pd.log

ipv4 = true
ipv6 = true

[precomputation]
elgamal = true

[upnp]
enabled = true
name = goI2Pd

[reseed]
verify = true

[addressbook]
subscriptions = http://inr.i2p/export/alive-hosts.txt,http://identiguy.i2p/hosts.txt,http://stats.i2p/cgi-bin/newhosts.txt,http://i2p-projekt.i2p/hosts.txt

### REASONING FOR CHANGING DEFAULT CONSOLE PORT
## We want to co-exist with other router projects peacefully inluding those that are on the same machine. This is a UI
## improvement project, not a router improvement project, and as such we will allow the use of any underlying I2P router.
[http]
enabled = true
address = 127.0.0.1
port = 7472

### REASONING FOR CHANGING DEFAULT HTTP PROXY PORT and DISABLING HTTP PROXY
## We want to co-exist with other router projects peacefully inluding those that are on the same machine. Disabling is
## the primary method of deferring to the parent router's SOCKS port. We change it in the example in case people using
## the embedded router want to re-enable it.
[httpproxy]
enabled = false
#address = 127.0.0.1
#port = 4454

### REASONING FOR CHANGING DEFAULT SOCKS PROXY PORT and DISABLING SOCKS PROXY
## We want to co-exist with other router projects peacefully inluding those that are on the same machine. Disabling is
## the primary method of deferring to the parent router's SOCKS port. We change it in the example in case people using
## the embedded router want to re-enable it.
[socksproxy]
enabled = false
#address = 127.0.0.1
#port = 4457

### REASONING FOR NOT CHANGING DEFAULT SAM PORT
## SAM clients do not normally require a password, so if a SAM hanshake can happen, it will more ofthen than not,
## succeed, causing the bundle start to fail if SAM is available. Regardless, we need a SAM on this port to run, leaving
## it here leaves us router-agnostic.
[sam]
enabled = true
address = 127.0.0.1
port = 7656

### REASONING FOR ENABLING I2PCONTROL and CHANGING DEFAULT I2PCONTROL PORT
## Java I2P provides I2PControl via an embedded application which is available under localhost 7657, whereas I2Pd
## makes it available at localhost:7650 as it's own service. In order to resolve this, our application will need to
## proxy requests to the underlying router or probe for it anyway. The port can't matter to the client behind the proxy
## probe thingy, but we need to guarantee i2pcontrol if we need to start an embedded router because of a SAM
## availability issue 
[i2pcontrol]
enabled = true
address = 127.0.0.1
port = 4450
#password = itoopie
`

var tunnelFile = `#
# tunnels.conf file intentionally left blank
#`

// Set the environment variable I2P_DIRECTORY_PATH to override the path returned by UnpackI2PdDir
var I2P_DIRECTORY_PATH = ""

// Returns nil if a file exists and an error for everything else. Used to check for file existence.
func FileOK(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return err
	} else {
		return err
	}
}

var walkFn = func(path string, fi os.FileInfo, r io.ReadSeeker, err error) error {
	if err != nil {
		log.Printf("can't stat file %s: %v\n", path, err)
		return nil
	}
	fmt.Println(path)
	if !fi.IsDir() {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Printf("can't read file %s: %v\n", path, err)
			return err
		}
		dir, err := UnpackI2PdDir()
		if err != nil {
			log.Printf("can't find path: %v\n", err)
			return err
		}
		err = ioutil.WriteFile(filepath.Join(dir, path), b, fi.Mode())
		if err != nil {
			log.Printf("can't write file %s: %v\n", filepath.Join(dir, path), err)
			return err
		}
		dirpath := strings.Split(path, ".")[0]
		log.Printf("wrote file %s: %v", filepath.Join(dir, path), fi.Mode())
		err = archiver.Unarchive(filepath.Join(dir, path), filepath.Join(dir, dirpath))
		if err != nil {
			log.Printf("can't unarchive file %s: %v\n", filepath.Join(dir, path), err)
			return err
		}
		log.Printf("unpacked file %s", filepath.Join(dir, path))
	}
	return nil
}

//WriteConfOptions generates a default config file for the bundle
func WriteConfOptions(targetdir string) error {
	if FileOK(filepath.Join(filepath.Dir(targetdir), "i2pd.conf")) != nil {
		err := ioutil.WriteFile(filepath.Join(filepath.Dir(targetdir), "i2pd.conf"), []byte(configFile), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

//WritetunnelOptions generates a default tunnel config file for the bundle
func WriteTunnelOptions(targetdir string) error {
	if FileOK(filepath.Join(filepath.Dir(targetdir), "tunnels.conf")) != nil {
		err := ioutil.WriteFile(filepath.Join(filepath.Dir(targetdir), "tunnels.conf"), []byte(tunnelFile), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

//WriteAllFiles generates an I2Pd installation and configuration for the bundle
func WriteAllFiles(targetdir string) error {
	if err := vfsutil.WalkFiles(FS, "/", walkFn); err != nil {
		return err
	}
	if err := WriteConfOptions(targetdir); err != nil {
		return err
	}
	if err := WriteTunnelOptions(targetdir); err != nil {
		return err
	}
	return nil
}

//
func UnpackI2PdPath() (string, error) {
	dirPath, err := UnpackI2PdDir()
	if err != nil {
		return "", err
	}
	ri2pd := filepath.Join(dirPath, "i2pd")
	return ri2pd, nil
}

//
func UnpackI2PdLibPath() (string, error) {
	dirPath, err := UnpackI2PdDir()
	if err != nil {
		return "", err
	}
	rlib := filepath.Join(dirPath, "lib")
	return rlib, nil
}

//
func UnpackI2PdDir() (string, error) {
	if I2P_DIRECTORY_PATH != "" {
		return I2P_DIRECTORY_PATH, nil
	}
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	ri2pd := filepath.Dir(executablePath)
	return ri2pd, nil
}
