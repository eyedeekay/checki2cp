package zerobundle

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var cmd *exec.Cmd

func GetZeroPID() int {
	return cmd.Process.Pid
}

func GetZeroProcess() *os.Process {
	return cmd.Process
}

func LatestZero() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(LatestZeroBinDir(), "i2p-zero.exe")
	} else {
		return filepath.Join(LatestZeroBinDir(), "i2p-zero")
	}
}

func LatestZeroJavaHome() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(LatestZeroBinDirJavaHome(), "i2p-zero.exe")
	} else {
		return filepath.Join(LatestZeroBinDirJavaHome(), "i2p-zero")
	}
}

func LatestZeroBinDir() string {
	var dir string
	var err error
	if dir, err = UnpackZeroDir(); err == nil {
		fs, er := ioutil.ReadDir(dir)
		if er != nil {
			log.Fatal(er)
		}
		if runtime.GOOS == "windows" {
			return filepath.Join(dir, fs[len(fs)-1].Name(), "router")
		} else {
			return filepath.Join(dir, fs[len(fs)-1].Name(), "router", "bin")
		}
	} else {
		log.Fatal(err)
	}
	return ""
}

func LatestZeroBinDirJavaHome() string {
	fs, er := ioutil.ReadDir(JAVA_I2P_OPT_DIR)
	if er != nil {
		log.Fatal(er)
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(JAVA_I2P_OPT_DIR, fs[len(fs)-1].Name(), "router")
	} else {
		return filepath.Join(JAVA_I2P_OPT_DIR, fs[len(fs)-1].Name(), "router", "bin")
	}
}

func StopZero() {
	if runtime.GOOS == "windows" {
		GetZeroProcess().Signal(os.Kill)
	} else {
		GetZeroProcess().Signal(os.Interrupt)
	}
}

func CommandZero() (*exec.Cmd, error) {
	if err := UnpackZero(); err != nil {
		log.Println(err)
	}
	latest := LatestZero()
	return exec.Command(latest), nil
}

func CommandZeroContext(ctx context.Context) (*exec.Cmd, error) {
	if err := UnpackZero(); err != nil {
		log.Println(err)
	}
	latest := LatestZero()
	return exec.CommandContext(ctx, latest), nil
}

func RunZero() error {
	var err error
	cmd, err = CommandZero()
	if err != nil {
		return err
	}
	return cmd.Run()
}

func StartZero() error {
	var err error
	cmd, err = CommandZero()
	if err != nil {
		return err
	}
	return cmd.Start()
}

func CommandZeroJavaHome() (*exec.Cmd, error) {
	if err := UnpackZeroJavaHome(); err != nil {
		log.Println(err)
	}
	latest := LatestZeroJavaHome()
	return exec.Command(latest), nil
}

func CommandZeroJavaHomeContext(ctx context.Context) (*exec.Cmd, error) {
	if err := UnpackZeroJavaHome(); err != nil {
		log.Println(err)
	}
	latest := LatestZeroJavaHome()
	return exec.CommandContext(ctx, latest), nil
}

func RunZeroJavaHome() error {
	var err error
	cmd, err = CommandZeroJavaHome()
	if err != nil {
		return err
	}
	return cmd.Run()
}

func StartZeroJavaHome() error {
	var err error
	cmd, err = CommandZeroJavaHome()
	if err != nil {
		return err
	}
	return cmd.Start()
}

func SAM() error {
	tcp, err := net.DialTCP("ip4", nil, &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8051})
	if err != nil {
		return err
	}
	defer tcp.Close()
	if runtime.GOOS == "windows" {
		tcp.Write([]byte("sam.create\r\n"))
	} else {
		tcp.Write([]byte("sam.create\n"))
	}
	return nil
}
