package checki2pcontrol

import (
	"github.com/eyedeekay/go-i2pcontrol"
)

func CheckI2PControlEcho(host, port, password string) (bool, error) {
	if host == "" {
	}
	if port == "" {
		port = "7650"
	}
	if password == "" {
		password = "itoopie"
	}
	i2pcontrol.Initialize(host, port, "jsonrpc")
	if _, err := i2pcontrol.Authenticate(password); err != nil {
		return false, err
	}
	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		return false, err
	}
	return true, nil
}
