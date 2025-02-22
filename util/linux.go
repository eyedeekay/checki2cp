//go:build linux
// +build linux

// Description: Linux-specific utility functions for detecting headless systems.
package util

import "os"

func IsHeadless() (bool, error) {
	// Method 1: Check common environment variables
	if os.Getenv("DISPLAY") != "" || os.Getenv("WAYLAND_DISPLAY") != "" {
		return false, nil
	}
	// Method 2: Check if running in SSH session
	if os.Getenv("SSH_TTY") != "" || os.Getenv("SSH_CONNECTION") != "" {
		return true, nil
	}
	return true, nil
}
