package checki2pcontrol

import "testing"

func TestCheckI2PControlEcho(t *testing.T) {
	if works, err := CheckI2PControlEcho("", "", ""); works {
		t.Log("Proxy success")
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Fatal("Proxy not found")
	}
}
