//+build generate

package main

import (
	"github.com/zserge/lorca"

	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var unpacker = `package REPLACEME

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func userFind() string {
	if os.Geteuid() == 0 {
		log.Fatal("Do not run this application as root!")
	}
	if un, err := os.UserHomeDir(); err == nil {
		os.MkdirAll(filepath.Join(un, "i2p"), 0755)
		return un
	}
	return ""
}

var userdir = filepath.Join(userFind(), "/i2p/opt/i2p-zero")

func writeFile(val os.FileInfo, system *fs) ([]byte, error) {
	if !val.IsDir() {
		file, err := system.Open(val.Name())
		if err != nil {
			return nil, err
		}
		sys := bytes.NewBuffer(nil)
		if _, err := io.Copy(sys, file); err != nil {
			return nil, err
		} else {
			return sys.Bytes(), nil
		}
	} else {
		log.Println(filepath.Join(userdir, val.Name()), "ignored", "contents", val.Sys())
	}
	return nil, fmt.Errorf("undefined unpacker error")
}

func WriteBrowser(FS *fs) ([]byte, error) {
	if embedded, err := FS.Readdir(-1); err != nil {
		log.Fatal("Extension error, embedded extension not read.", err)
	} else {
		for _, val := range embedded {
			if val.IsDir() {
				os.MkdirAll(filepath.Join(userdir, val.Name()), val.Mode())
			} else {
				return writeFile(val, FS)
			}
		}
	}
	return nil, nil
}
`

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.

	if err := deleteDirectories(); err != nil {
		log.Fatal(err)
	}
	if err := createDirectories(); err != nil {
		log.Fatal(err)
	}
	if err := generateGoUnpacker(); err != nil {
		log.Fatal(err)
	}
	if err := splitBinaries("i2p-zero-linux.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
	if err := updateAllChunks("linux", "i2p-zero-linux.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
	if err := splitBinaries("i2p-zero-win.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
	if err := updateAllChunks("windows", "i2p-zero-win.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
	if err := splitBinaries("i2p-zero-darwin.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
	if err := updateAllChunks("darwin", "i2p-zero-darwin.v1.17.zip"); err != nil {
		log.Fatal(err)
	}
}

var libs = []string{
	"aa",
	"ab",
	"ac",
	"ad",
	"ae",
	"af",
	"ag",
	"ah",
	"ai",
	"aj",
	"ak",
	"al",
}

func updateChunk(chunk, platform, file string) error {
	err := lorca.Embed("iz"+chunk, "parts/"+chunk+"/chunk_"+platform+".go", file+"."+chunk)
	if err != nil {
		return err
	}
	log.Println("embedded iz" + chunk)
	return nil
}

func updateAllChunks(platform, file string) error {
	for _, lib := range libs {
		updateChunk(lib, platform, file)
	}
	return nil
}

func splitBinaries(fileToBeChunked string) error {
	bytes, err := ioutil.ReadFile(fileToBeChunked)
	if err != nil {
		return err
	}
	chunkSize := len(bytes) / 12
	for index, lib := range libs {
		start := index * chunkSize
		finish := ((index + 1) * chunkSize)
		if index == 11 {
			finish = len(bytes)
		}
		outBytes := bytes[start:finish]
		err := ioutil.WriteFile(fileToBeChunked+"."+lib, outBytes, 0644)
		if err != nil {
			return err
		}
		log.Printf("Started at: %d,  Ended at: %d", start, finish)
	}
	return nil
}

func deleteDirectories() error {
	for _, dir := range libs {
		err := os.RemoveAll(filepath.Join("parts", dir))
		if err != nil {
			return err
		}
	}
	return nil
}

func createDirectories() error {
	for _, dir := range libs {
		err := os.MkdirAll(filepath.Join("parts", dir), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateGoUnpacker() error {
	for index, dir := range libs {
		contents := strings.Replace(unpacker, "REPLACEME", "iz"+libs[index], -1)
		if err := ioutil.WriteFile(filepath.Join("parts", dir, "unpacker.go"), []byte(contents), 0644); err != nil {
			return err
		}
	}
	return nil
}
