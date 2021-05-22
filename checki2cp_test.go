package checki2p

import (
	"testing"
)

func TestRouterLaunch(t *testing.T) {
	ok, err := ConditionallyLaunchI2P()
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Log("I2P is installed, successfully confirmed", ok)
	} else {
		t.Log("I2P is in a default location, user feedback is needed")
	}
}

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
		t.Log("I2P is installed, successfully confirmed", ok)
	} else {
		t.Log("I2P is in a default location, user feedback is needed")
	}
}

func TestRouterPath(t *testing.T) {
	ok, err := FindI2PIsInstalledDefaultLocation()
	if err != nil {
		t.Fatal(err)
	}
	if ok != "" {
		t.Log("I2P is installed, successfully confirmed", ok)
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
