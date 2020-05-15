package checksam

import (
	"github.com/eyedeekay/sam3"
)

// CheckSAMAvailable tries a SAM connection and returns true if it succeeds.
func CheckSAMAvailable(yoursam string) bool {
	if yoursam == "" {
		yoursam = "127.0.0.1:7656"
	}
	if sam, err := sam3.NewSAM(yoursam); err != nil {
		return false
	} else {
		defer sam.Close()
		if _, err := sam.NewKeys(); err != nil {
			return false
		}
	}
	return true
}
