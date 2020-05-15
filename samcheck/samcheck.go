package checksam

import (
	"github.com/eyedeekay/sam3"
)

// CheckSAMAvailable tries a SAM connection and returns true if it succeeds.
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
