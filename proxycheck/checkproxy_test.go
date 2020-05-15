package checkproxy

import "testing"

func TestProxyDotI2P(t *testing.T) {
	if ProxyDotI2P() {
		t.Log("Proxy success")
	} else {
		t.Fatal("Proxy not found")
	}
}
