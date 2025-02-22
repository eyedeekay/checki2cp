package util

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Inithome returns the path to the user's home directory
// with the given string appended to the end.
// It panics if the user's home directory cannot be found.
func Inithome(str string) string {
	s, e := os.UserHomeDir()
	if e != nil {
		panic(e)
	}
	path := filepath.Join(s, str)
	log.Println(path)
	return path
}

// CheckFileExists checks to see if a file exists at the given path
// and returns a boolean value.
// It returns true if the file exists, false otherwise.
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Home returns the path to the I2P(Java) home directory
// based on the current operating system.
// It returns a string containing the path to the I2P home directory.
func Home() string {
	if runtime.GOOS == "windows" {
		return "\\i2p"
	} else if runtime.GOOS == "darwin" {
		return "/Library/Application Support/i2p"
	}
	return "/.i2p"
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
