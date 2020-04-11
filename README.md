checki2cp
=========

Library and terminal application which checks for the presence of a usable i2p
router by attempting to connect to i2cp. Yet another tiny but essential function
I don't want to be bothered with writing a billion times.

Well like most things, it grew in complexity. Now it does about 5 or 6 things,
it determines presence, running state, path, router "Style"(Java I2P, i2pd, 
Zero) and it can start(but not stop) an I2P router. If you want to stop it,
use i2pcontrol.




I2PD Embedding Tools:
---------------------



### Roadmap

 1. Do a 100% static build of our own, less getHostByName peculiarities.
 2. Set up router with SWIG instead of with static embedded executable.
 3. Minimize number of components we build(nothing but I2PControl, SAM, and
  I2CP).
 4. Provide easy router/client setup functions to import i2p routers into
  applications with minimal boilerplate.
 3. Pure-Go I2PControl implementation.
 4. Pure-Go SAM implementation.
 5. Pure-Go I2CP implementation.

Here's where things get a little wierder.
-----------------------------------------

Have you ever wondered whether it's possible to compile a 200-ish Megabyte .go
file? Because before today, I hadn't. Turns out, it's looking... kinda possible.
But it's a sonafagun to actually accomplish. Let me back up, today, I wanted to
see if I could bundle I2P-Zero inside of a Go package wholesale. To do this, I
went ahead and used lorca's embedding code to do so, because that's what I've
been using for other things that this will be used in. Also it's easy to use.
That means it takes the whole, compiled I2P-Zero package(zip file or whatever)
encodes it as a string, and puts it in an object inside a Go file. Those, you
can find on the releases page, under releases starting with z, followed by the
corresponding Java I2P Major and Minor version, followed by the i2cpcheck
version. So:

        z9.46.12

for example.

Download all 3 of those to the zerobundle directory and compile them. You now
have a virtual filesystem, containing the I2P Zero application as a zip file,
which you can unpack to the disk at a desirable install location. Convenience
functions to pack and unpack the data will be provided when I am done evaluating
non-lorca options that might achieve smaller files, since I know I can achieve
much smaller files with i2pd anyway.

**TL:DR Don't use this part unless you want to give yourself a headache.** It's 
not imported by anything else, use the i2pd version instead and if Zero becomes 
a realistic possibility in the future I'll change this section to reflect that.