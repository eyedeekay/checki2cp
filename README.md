checki2cp
=========

[![Go Report Card](https://goreportcard.com/badge/github.com/eyedeekay/checki2cp)](https://goreportcard.com/report/github.com/eyedeekay/checki2cp)
[![Documentation](https://godoc.org/github.com/eyedeekay/checki2cp?status.svg)](http://godoc.org/github.com/eyedeekay/checki2cp)

Library and terminal application which checks for the presence of a usable i2p
router by attempting various I2P client-related probes and tests. Includes 
everything you need to completely embed I2Pd in a Go application on Linux,
OSX, and Windows, or use them as modular tools for checking the status of your
I2P router.

Directories
-----------

 * **./**  **-** Files in the base directory contain functions that check for the presence of an I2P router in a default
  location, or for the presence of an I2CP port. Makefile is also here.
 * **./controlcheck**  **-** Detects whether an I2PControl API endpoint is available, and provides a tool for finding
  one if it is available. It is not recommended to use I2PControl checking for presence detection, you will need to
  parse the errors to know what is going on.
 * **./i2cpcheck**  **-** A terminal application which probes for an I2P router by looking in a default location, the
  $PATH, or by probing I2CP.
 * **./i2pdbundle**  **-** A set of tools and libraries for embedding an I2Pd router in a Go application, then
  installing it to a location under the application's control.
 * **./proxycheck**  **-** A tool for determining the presence of an I2P router by making a request(to "proxy.i2p" by
  default) over an I2P HTTP Proxy.
 * **./samcheck**  **-** A tool for determining the presence of an I2P router by doing a brief interaction with the SAM
  API.

I2P Router Presence Detection tools
-----------------------------------

Currently the command-line tool only does presence detection by checking for I2CP and or an I2P router installed in a
default location.

### Examples:

#### ./ Base Directory

##### Checking for an I2P Router by default install location

```Go
ok, err := CheckI2PIsInstalledDefaultLocation()
if err != nil {
    t.Fatal(err)
}
if ok {
    t.Log("I2P is installed, successfully confirmed")
} else {
    t.Log("I2P is in a default location, user feedback is needed")
}
```

##### Checking for an I2P Router by I2CP Port

```Go
ok, err := CheckI2PIsRunning()
if err != nil {
    t.Fatal(err)
}
if ok {
    t.Log("I2P is running, successfully confirmed I2CP")
} else {
    t.Log("I2P is not running, further testing is needed")
}
```

##### Launching an installed I2P router from the Library

TODO: Make this function work better with i2pd, find a way to integrate it into into the tests, then write the example.

#### ./proxycheck

##### Make a request through the default I2P HTTP Proxy to test presence

```Go
if ProxyDotI2P() {
    t.Log("Proxy success")
} else {
    t.Fatal("Proxy not found")
}
```

##### Use a non-default proxy instead

It honors the ```http_proxy``` environment variable, so just set it(I like to set both of these in case some system is
weird, I suppose it's a meager measure.):

```Go
if err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:4444"); err != nil {
    return false
}
if err := os.Setenv("http_proxy", "http://127.0.0.1:4444"); err != nil {
    return false
}
if ProxyDotI2P() {
    t.Log("Proxy success")
} else {
    t.Fatal("Proxy not found")
}
```

#### ./samcheck 

##### Check if SAM is available on a default port

```Go
if CheckSAMAvailable("") {
    t.Log("SAM success")
} else {
    t.Fatal("SAM not found")
}
```

##### Check if SAM is available on a non-default host or port

```Go
if CheckSAMAvailable("127.0.1.1:1234") {
    t.Log("SAM success")
} else {
    t.Fatal("SAM not found")
}
```

I2PD Embedding Tools:
---------------------

In the very far future, it would be cool to have a 100% pure-Go I2P router, but for right now, what I want to work on is
a way of working with I2P and with Go applications, **without** requiring undue knowledge on the part of the user. The
theory is that they want to install one application at a time, and don't want to run more than one I2P router unless
they need to. So the embedding tools assume that if they find an I2P router, that they should use that I2P router. At
this time, almost any useful I2P configuration will be detected and the embedded router will not start. In the future,
this behavior will be configurable.

#### ./controlcheck

##### See if an I2PControl API is present on any default port

```Go
if itworks, err := CheckI2PControlEcho("", "", "", ""); works {
    t.Log("Proxy success")
} else if err != nil {
    t.Fatal(err)
} else {
    t.Fatal("Proxy not found")
}
```

##### Check what host:port/path corresponds to an available default I2PControl API

TODO: Explain why you need this if your goal is to tolerate Java, I2Pd, and Embedded I2Pd all at once

```Go
if host, port, path, err := GetDefaultI2PControlPath(); err != nil {
    t.Fatal("I2PControl Not found")
} else {
    t.Log("I2Pcontrol found at", host, port, path)
}
```
