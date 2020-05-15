package checki2pcontrol

import (
	"github.com/eyedeekay/go-i2pcontrol"
)

func CheckI2PControlEcho(host, port, password, path string) (bool, error) {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "7650"
	}
	if password == "" {
		password = "itoopie"
	}
	i2pcontrol.Initialize(host, port, path)
	var finalError error
	if _, err := i2pcontrol.Authenticate(password); err != nil {
		finalError = err
	}
	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		finalError = err
	}
	if finalError == nil {
		return true, nil
	}
	finalError = nil
	port = "7657"
	password = "itoopie"
	path = "jsonrpc"
	i2pcontrol.Initialize(host, port, path)
	if _, err := i2pcontrol.Authenticate(password); err != nil {
		finalError = err
	}
	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		finalError = err
	}
	if finalError == nil {
		return true, nil
	}
	return false, finalError
}

func GetDefaultI2PControlPath() (string, string, string, error) {
	host := "127.0.0.1"
	port := "7650"
	password := "itoopie"
	path := ""
	i2pcontrol.Initialize(host, port, path)
	var finalError error
	if _, err := i2pcontrol.Authenticate(password); err != nil {
		finalError = err
	}
	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		finalError = err
	}
	if finalError == nil {
		return host, port, path, nil
	}
	finalError = nil
	port = "7657"
	password = "itoopie"
	path = "jsonrpc"
	i2pcontrol.Initialize(host, port, path)
	if _, err := i2pcontrol.Authenticate(password); err != nil {
		finalError = err
	}
	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		finalError = err
	}
	if finalError == nil {
		return host, port, path, nil
	}
	return "127.0.0.1", "4450", "", nil
}
