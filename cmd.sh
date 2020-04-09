BOOST_ROOT=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/boost \
    cmake -G "Unix Makefiles" \
    /media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/i2pd/build \
    -DBUILD_TYPE=Release \
    -DCMAKE_TOOLCHAIN_FILE=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/toolchain-mingw.cmake \
    -DWITH_AESNI=ON -DWITH_UPNP=ON -DWITH_STATIC=ON -DWITH_HARDENING=ON \
    -DCMAKE_INSTALL_PREFIX:PATH=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/i2pd/i2pd-mingw-64-static \
    -DZLIB_ROOT=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/stage \
    -DMINIUPNPC_LIBRARY=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/miniupnpc \
    -DBOOST_LIBRARYDIR:PATH=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/boost/stage-mingw-64/lib \
    -DOPENSSL_ROOT_DIR:PATH=/media/longterm/go/src/github.com/eyedeekay/checki2cp/i2pdbundle/stage
