//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	/*
		lorca.Embed("i2pd", "assets_linux_amd64.go", "linux")
		lorca.Embed("i2pd", "assets_windows_amd64.go", "win")
	*/

	lorca.Embed("i2pd", "assets_darwin.go", "mac")
	lorca.Embed("i2pd", "assets_linux_amd64.go", "linux_amd64")
	lorca.Embed("i2pd", "assets_windows_amd64.go", "win_amd64")
	/*	lorca.Embed("i2pd", "assets_linux_i386.go", "linux_i386")
		lorca.Embed("i2pd", "assets_windows_i386.go", "win_i386")
	*/
	lorca.Embed("i2pdtest", "test_files/test_assets.go", "test")
}
