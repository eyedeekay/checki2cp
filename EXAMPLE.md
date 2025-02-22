I2P Router Presence Detection tools
===================================

Currently the command-line tool only does presence detection by checking for I2CP and or an I2P router installed in a
default location.

## Examples:

### ./ Base Directory

#### Checking for an I2P Router by default install location

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

#### Checking for an I2P Router by I2CP Port

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

#### Launching an installed I2P router from the Library

TODO: Make this function work better with i2pd, find a way to integrate it into into the tests, then write the example.

### ./proxycheck

#### Make a request through the default I2P HTTP Proxy to test presence

```Go
if ProxyDotI2P() {
    t.Log("Proxy success")
} else {
    t.Fatal("Proxy not found")
}
```

#### Use a non-default proxy instead

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

### ./samcheck 

#### Check if SAM is available on a default port

```Go
if CheckSAMAvailable("") {
    t.Log("SAM success")
} else {
    t.Fatal("SAM not found")
}
```

#### Check if SAM is available on a non-default host or port

```Go
if CheckSAMAvailable("127.0.1.1:1234") {
    t.Log("SAM success")
} else {
    t.Fatal("SAM not found")
}
```

### ./controlcheck

#### See if an I2PControl API is present on any default port

```Go
if itworks, err := CheckI2PControlEcho("", "", "", ""); works {
    t.Log("Proxy success")
} else if err != nil {
    t.Fatal(err)
} else {
    t.Fatal("Proxy not found")
}
```

#### Check what host:port/path corresponds to an available default I2PControl API

TODO: Explain why you need this if your goal is to tolerate Java, I2Pd, and Embedded I2Pd all at once

```Go
if host, port, path, err := GetDefaultI2PControlPath(); err != nil {
    t.Fatal("I2PControl Not found")
} else {
    t.Log("I2Pcontrol found at", host, port, path)
}
``