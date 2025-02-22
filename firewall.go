package checki2p

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/eyedeekay/checki2cp/util"
)

// GetFirewallPort finds the configured UDP port of your I2P router to help
// configure firewalls. It does this by finding the router.config and reading
// it. This function will not work on I2Pd routers yet but it should be easy
// to add support once I get some more time to test and research it.
func GetFirewallPort() (string, error) {
	log.Println(util.I2P_ASUSER_HOME_LOCATION)
	file, err := ioutil.ReadFile(util.I2P_ASUSER_HOME_LOCATION + "/router.config")
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
