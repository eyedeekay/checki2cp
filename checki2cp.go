package checki2p

import (
	"fmt"
	"log"

	checksam "github.com/eyedeekay/checki2cp/samcheck"
	go_i2cp "github.com/go-i2p/go-i2cp"
)

func i2pdArgs() ([]string, error) {
	return []string{""}, nil
}

// CheckI2PIsRunning is frequently the only thing I need a reliable, non-SAM
// way to test at runtime.
func CheckI2PIsRunning() (bool, error) {
	log.Println("Trying to discover a running I2P router")
	if checksam.CheckSAMAvailable("127.0.0.1:7656") {
		log.Println("I2P is running with a SAM API exposed.")
		return true, nil
	}
	if val, err := CheckI2CPConnection(); err != nil {
		return false, nil
	} else if val {
		log.Println("I2P is running with an I2CP API exposed.")
		return true, nil
	}
	return false, nil
}

// CheckI2CPConnection is determines if I2CP is available on the TCP port
// usually used by I2P routers(7654).
func CheckI2CPConnection() (bool, error) {
	log.Println("Trying to discover a running I2P router")
	client := go_i2cp.NewClient(nil)
	err := client.Connect()
	if err != nil {
		return false, nil
	}
	destination, err := go_i2cp.NewDestination()
	if err != nil {
		return false, err
	}
	if destination == nil {
		return false, fmt.Errorf("Client connection was made but no destination could be generated.")
	}
	client.Disconnect()
	log.Println("I2P is running.")
	return true, nil
}
