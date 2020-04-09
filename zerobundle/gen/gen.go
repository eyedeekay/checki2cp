//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("zero", "linux/assets.go", "i2p-zero/dist-zip/i2p-zero-linux.v1.16.zip")
    lorca.Embed("zero", "mac/assets.go", "i2p-zero/dist-zip/i2p-zero-mac.v1.16.zip")
    lorca.Embed("zero", "windows/assets.go", "i2p-zero/dist-zip/i2p-zero-win.v1.16.zip")
}

