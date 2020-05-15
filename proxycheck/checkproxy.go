package checkproxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ProxyDotI2P makes a request to proxy.i2p. If HTTP_PROXY is unset, then it defaults to using localhost 4444
func ProxyDotI2P() bool {
	return ProxyGetSite("http://proxy.i2p")
}

// ProxyGetSite fetches a URL via the I2P HTTP Proxy. The desired site must include the http(s):// protocol prefix
func ProxyGetSite(url string) bool {
	if prox := os.Getenv("HTTP_PROXY"); prox == "" {
		if err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:4444"); err != nil {
			return false
		}
		if err := os.Setenv("http_proxy", "http://127.0.0.1:4444"); err != nil {
			return false
		}
	}
	// make a sample HTTP GET request
	// check for response error
	if res, err := http.Get(url); err != nil {
		return false
	} else {
		// close response body
		defer res.Body.Close()
		// read all response body
		if data, err := ioutil.ReadAll(res.Body); err != nil {
			return false
		} else {
			// print `data` as a string
			log.Printf("%s\n", data)
		}
	}
	return true
}
