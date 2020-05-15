package checksam

import (
	"github.com/eyedeekay/sam3"
)

//
func CheckSAMAvailable(yoursam string) bool {
	if yoursam == "" {
		yoursam = "127.0.0.1:7656"
	}
	if sam, err := sam3.NewSAM(yoursam); err != nil {
		return false
	} else {
		if _, err := sam.NewKeys(); err != nil {
			return false
		} else {
			return true
		}
	}
	return false
}
