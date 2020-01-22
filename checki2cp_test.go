package checki2p

import (
	"testing"
)

func TestClient(t *testing.T) {
	ok, err := CheckI2PIsRunning()
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Log("I2P is running, successfully confirmed I2CP")
	} else {
		t.Log("I2P is not running, further testing is needed")
	}
}

func TestRouter(t *testing.T) {
	ok, err := CheckI2PIsInstalledDefaultLocation()
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Log("I2P is installed, successfully confirmed")
	} else {
		t.Log("I2P is in a default location, user feedback is needed")
	}
}

func TestFirewallPort(t *testing.T) {
	port, err := GetFirewallPort()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Port was", port)
}
