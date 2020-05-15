package checki2pcontrol

import "testing"

func TestCheckI2PControlEcho(t *testing.T) {
	if works, err := CheckI2PControlEcho("", "", "", ""); works {
		t.Log("Proxy success")
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Fatal("Proxy not found")
	}
	if host, port, path, err := GetDefaultI2PControlPath(); err != nil {
		t.Fatal("I2PControl Not found")
	} else {
		t.Log("I2Pcontrol found at", host, port, path)
	}
}
