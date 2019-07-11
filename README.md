# checki2cp

Library and terminal application which checks for the presence of a usable i2p
router by attempting to connect to i2cp. Yet another tiny but essential function
I don't want to be bothered with writing a billion times.

It contains just two functions, for helping Go applications confirm the presence
of an I2P router on the system. The first is *CheckI2PIsRunning*, which attempts
to connect to the I2CP Port, generate a destination, and quit. If this is
successful, an i2p router is obviously installed. The second is
*CheckI2PIsInstalledDefaultLocation* which checks the default locations for
the I2P router on various platforms to confirm whether an i2p router is present.

It does not launch the router or anything else to do with managing the router.
It is solely for checking whether a router is present.
