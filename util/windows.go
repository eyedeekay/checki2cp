//go:build windows
// +build windows

// Description: Windows-specific utility functions for detecting headless systems.
// I've never actually used these on a headless Windows system, so I'm not sure if they work.
package util

import (
	"os"
	"syscall"
	"unsafe"
)

// Windows API constants
const (
	// SM_REMOTESESSION (0x1000) indicates if the current session is remote
	SM_REMOTESESSION = 0x1000

	// Invalid session ID constant used by WTSGetActiveConsoleSessionId
	INVALID_SESSION_ID = 0xFFFFFFFF
)

// IsHeadless determines if the current Windows system is running in a headless
// environment (i.e., without a display/GUI) using multiple detection methods.
//
// Returns:
//   - bool: true if the system appears to be headless, false otherwise
//   - error: any error encountered during the detection process
//
// Detection methods:
//  1. Checks NO_DISPLAY environment variable
//  2. Verifies the active console session ID
//  3. Queries Windows UI metrics for remote session status
func IsHeadless() (bool, error) {
	// Method 1: Check environment variable
	// Some headless systems explicitly set this variable
	if os.Getenv("NO_DISPLAY") == "1" {
		return true, nil
	}

	// Method 2: Check active console session
	// Load WTS API functions dynamically
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getActiveSessionProc := kernel32.NewProc("WTSGetActiveConsoleSessionId")

	sessionId, _, _ := getActiveSessionProc.Call()
	if sessionId == INVALID_SESSION_ID {
		return true, nil
	}

	// Method 3: Check for an active display device
	displayAvailable, err := IsDisplayAvailable()
	if err != nil {
		return true, err
	}
	if displayAvailable {
		return false, nil
	}

	// Method 3: Check system metrics for remote session
	// Load User32 API functions dynamically
	user32 := syscall.NewLazyDLL("user32.dll")
	getSystemMetricsProc := user32.NewProc("GetSystemMetrics")

	// Query if we're in a remote session
	ret, _, err := getSystemMetricsProc.Call(uintptr(SM_REMOTESESSION))
	if err != nil {
		return true, err
	}

	return ret != 0, nil
}

// IsDisplayAvailable checks specifically if a display device is available
// by attempting to enumerate display devices.
//
// Returns:
//   - bool: true if at least one display device is available, false otherwise
//   - error: any error encountered during the check
func IsDisplayAvailable() (bool, error) {
	user32 := syscall.NewLazyDLL("user32.dll")
	enumDisplayDevicesProc := user32.NewProc("EnumDisplayDevicesW")

	type DISPLAY_DEVICE struct {
		Cb           uint32
		DeviceName   [32]uint16
		DeviceString [128]uint16
		StateFlags   uint32
		DeviceID     [128]uint16
		DeviceKey    [128]uint16
	}

	var dd DISPLAY_DEVICE
	dd.Cb = uint32(unsafe.Sizeof(dd))

	// Try to get information about the primary display adapter (device index 0)
	ret, _, err := enumDisplayDevicesProc.Call(
		0,                            // lpDevice (NULL for primary)
		0,                            // iDevNum (device index)
		uintptr(unsafe.Pointer(&dd)), // lpDisplayDevice
		0,                            // dwFlags
	)
	if err != nil {
		return false, err
	}

	return ret != 0, nil
}
