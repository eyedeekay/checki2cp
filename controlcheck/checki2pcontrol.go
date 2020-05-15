package checki2pcontrol

import (
	"github.com/eyedeekay/go-i2pcontrol"
)

// CheckI2PControlEcho attempts a connection and an echo command on it.
// it returns true if the command is successful and false, with an error,
// if not.
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

// GetDefaultI2PControlPath probes default locations for the I2PControl API, returning
// either a working I2PControl API and no error, or the defaults of the embedded router
// and an error
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
