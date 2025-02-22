package constant

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
	log.Println(filepath.Join(s, str))
	return filepath.Join(s, str)
}

func CheckFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func Home() string {
	if runtime.GOOS == "windows" {
		return "\\i2p"
	} else if runtime.GOOS == "darwin" {
		return "/Library/Application Support/i2p"
	}
	return "/.i2p"
}
