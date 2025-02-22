package util

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Inithome(str string) string {
	s, e := os.UserHomeDir()
	if e != nil {
		panic(e)
	}
	path := filepath.Join(s, str)
	log.Println(path)
	return path
}

func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Home() string {
	if runtime.GOOS == "windows" {
		return "\\i2p"
	} else if runtime.GOOS == "darwin" {
		return "/Library/Application Support/i2p"
	}
	return ".i2p"
}

// CheckHeadlessness checks to see if the current environment is headless
// and returns a boolean value.
// it calls a platform-specific function in the background.
// it returns true if the system appears to be headless, false otherwise
// and any error encountered during detection.
func CheckHeadlessness() (bool, error) {
	switch runtime.GOOS {
	case "windows":
		return IsHeadless()
	case "darwin":
		return IsHeadless()
	case "linux":
		return IsHeadless()
	}
	return false, nil
}
