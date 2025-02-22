//go:build darwin
// +build darwin

// Description: Darwin-specific utility functions for detecting headless systems.
// I've never actually used these on a headless macOS system, so I'm not sure if they work.
// But I'm slightly surer than the Windows ones.
package util

import (
	"os"
	"os/exec"
	"strings"
)

// IsHeadless determines if the current macOS system is running headless
// using multiple detection methods.
//
// Returns:
//   - bool: true if the system appears to be headless, false otherwise
//   - error: any error encountered during detection
func IsHeadless() (bool, error) {
	// Method 1: Check common environment variables
	if os.Getenv("DISPLAY") == "" {
		return true, nil
	}

	// Method 2: Check if running in SSH session
	if os.Getenv("SSH_TTY") != "" || os.Getenv("SSH_CONNECTION") != "" {
		return true, nil
	}

	return IsDisplayAvailableMacOS()
}

// IsDisplayAvailableMacOS checks specifically for connected displays
// using ioreg command, which is more reliable than system_profiler
// for detecting actual display hardware.
//
// Returns:
//   - bool: true if at least one display is connected, false otherwise
//   - error: any error encountered during the check
func IsDisplayAvailableMacOS() (bool, error) {
	// Check for connected displays using ioreg
	cmd := exec.Command("ioreg", "-l", "-w", "0", "-r", "-c", "IODisplayConnect")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// If any display is connected, the output will contain "IODisplayConnect"
	return strings.Contains(string(output), "IODisplayConnect"), nil
}
