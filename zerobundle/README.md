Easy I2P Zero Bundling Tool for Go
==================================

This accomplishes the bundling and managment of an I2P Zero router from within a Go application. The recommended use of
this library is as a way of providing an embedded *option* in an out-of-tree application, an out-of-tree application
which would otherwise prefer an existing I2P Router. This means that the standard way to use it should be:

1. Check for a running I2P Router. If one is found, skip to step *5.*

2. Check for a stopped I2P Router installed from an official package.

3. If a stopped package-installed I2P router is available from an official package, start it. Skip step *4.*

4. If no other I2P Router is available, start the embedded I2P Zero Router.

5. Start the external application.

Doing things in this way allows us to conserve resources by not running redundant I2P Routers on the same computer,
while also allowing the use of an embedded I2P router to auto-configure standalone applications on computer even when
I2P is not present.

Use Scenarios
-------------

**ABSOLUTELY do not** use it for applications that use i2ptunnel via ```tunnel-control.sh```! **SAM only**. A goal of
this is to build apps that can seamlessly switch between an *embedded* I2P Router provided by this Go package and a
*package-provided* I2P Router which may be installed per-user or system-wide as the case may be. The best way to achieve
this is with SAM.

### **Scenario A:** I2P Router is installed on host PC *Prior to* the first run of the out-of-tree application.

In this scenario, the I2P Router in use is the package-installed router, and the embedded one is left alone.

### **Scenario B:** out-of-tree application is installed on host PC *alone*, with no other router available

In this scenario, the I2P Router in use is the embedded router.

### **Scenario C:** out-of-tree application is installed on host PC *prior to* a system-wide I2P router which becomes preferred.

In this scenario, the keys used for identifying the SAM tunnels are managed by the application and thus migrate with
the application from the embedded router to the package-installed router.

### **Scenario D:** Misuse, tunnel-control.sh tunnels are somehow used

More details on this scenario later. Don't do it. This might change sometime in the future but probably not soon.

Example Usage:
--------------

        package main

        import (
          "log"
        )

        import (
          "github.com/eyedeekay/checki2cp/zerobundle"
          "github.com/eyedeekay/checki2cp"
        )

        func main() {
          if ok, err := checki2p.ConditionallyLaunchI2P(); ok {
            if err != nil {
              log.Println(err)
            }
          } else {
            if err := zerobundle.UnpackZero(); err != nil {
              log.Println(err)
            }
            latest := zerobundle.LatestZero()
            log.Println("latest zero version is:", latest)
            if err := zerobundle.StartZero(); err != nil {
              log.Fatal(err)
            }
            log.Println("Starting SAM")
            if err := zerobundle.SAM(); err != nil {
              log.Fatal(err)
            }
            log.Println("Undefined I2P launching error")
          }
        }

