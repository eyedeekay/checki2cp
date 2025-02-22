I2P Router Presence Detection tools
===================================

Currently, the command-line tool performs presence detection by checking for I2CP and I2P router installations in default locations.

## Examples:

### ./ Base Directory

#### Checking for an I2P Router by default install location

```Go
ok, err := CheckI2PIsInstalledDefaultLocation()
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("I2P is installed, successfully confirmed")
} else {
    log.Println("I2P is not in a default location")
}
```

#### Checking for an I2P Router by I2CP Port

```Go
ok, err := CheckI2PIsRunning()
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("I2P is running, successfully confirmed I2CP")
} else {
    log.Println("I2P is not running")
}
```

### ./proxycheck

#### Make a request through the default I2P HTTP Proxy to test presence

```Go
ok, err := ProxyDotI2P()
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("Proxy success")
} else {
    log.Fatal("Proxy not found")
}
```

#### Use a non-default proxy instead

It honors the `http_proxy` environment variable:

```Go
if err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:4444"); err != nil {
    log.Fatal(err)
}
ok, err := ProxyDotI2P()
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("Proxy success")
} else {
    log.Fatal("Proxy not found")
}
```

### ./samcheck 

#### Check if SAM is available on a default port

```Go
ok, err := CheckSAMAvailable("")
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("SAM success")
} else {
    log.Fatal("SAM not found")
}
```

#### Check if SAM is available on a non-default host or port

```Go
ok, err := CheckSAMAvailable("127.0.1.1:1234")
if err != nil {
    log.Fatal(err)
}
if ok {
    log.Println("SAM success")
} else {
    log.Fatal("SAM not found")
}
```

### ./controlcheck

#### See if an I2PControl API is present on any default port

```Go
works, err := CheckI2PControlEcho("", "", "", "")
if err != nil {
    log.Fatal(err)
}
if works {
    log.Println("I2PControl success")
} else {
    log.Fatal("I2PControl not found")
}
```

#### Check what host:port/path corresponds to an available default I2PControl API

```Go
host, port, path, err := GetDefaultI2PControlPath()
if err != nil {
    log.Fatal("I2PControl Not found")
}
log.Println("I2PControl found at", host, port, path)
```