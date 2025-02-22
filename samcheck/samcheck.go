package checksam

import (
	"github.com/go-i2p/sam3"
)

// CheckSAMAvailable tries a SAM connection and returns true if it succeeds.
// If yoursam is empty, it will try to connect to the default SAM address.
func CheckSAMAvailable(yoursam string) bool {
	if yoursam == "" {
		yoursam = "127.0.0.1:7656"
	}
	sam, err := sam3.NewSAM(yoursam)
	if err != nil {
		return false
	}
	defer sam.Close()
	if _, err := sam.NewKeys(); err != nil {
		return false
	}
	return true
}
