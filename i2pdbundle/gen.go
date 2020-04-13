//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("i2pd", "assets_linux.go", "linux")
    lorca.Embed("i2pd", "assets_darwin.go", "mac")
    lorca.Embed("i2pd", "assets_windows.go", "win")
}

