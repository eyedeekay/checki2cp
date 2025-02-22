checki2cp
=========

[![Go Report Card](https://goreportcard.com/badge/github.com/go-i2p/checki2cp)](https://goreportcard.com/report/github.com/go-i2p/checki2cp)
[![Documentation](https://godoc.org/github.com/go-i2p/checki2cp?status.svg)](http://godoc.org/github.com/go-i2p/checki2cp)

Library and terminal application which checks for the presence of a usable i2p
router by attempting various I2P client-related probes and tests. It includes 
everything you need to completely embed I2Pd in a Go application on Linux, 
OSX, and Windows. You can also use them as modular tools for checking the 
status of your I2P router.

Directories
-----------

- **./** - Files in the base directory contain functions that check for the presence of an I2P router in a default
  location. It also checks for the presence of an I2CP port. The Makefile is also here.
- **./controlcheck** - Detects whether an I2PControl API endpoint is available. It provides a tool for finding
  one if it is available. It is not recommended to use I2PControl checking for presence detection, you will need to
  parse the errors to know what is going on.
- **./i2cpcheck** - A terminal application which probes for an I2P router by looking in a default location. It also checks the
  $PATH, or by probing I2CP.
- **./i2pdbundle** - A set of tools and libraries for embedding an I2Pd router in a Go application.
  It then installs the router to a location under the application's control.
- **./proxycheck** - A tool for determining the presence of an I2P router. It does this by making a request to "proxy.i2p" by
  default) over an I2P HTTP Proxy.
- **./samcheck** - A tool for determining the presence of an I2P router. It does this by performing a brief interaction with the SAM
  API.
- **./util** - Various tools for helping manage I2P routers. It also includes tools for getting information from the host system.

