

GO111MODULE=on

VERSION=0.0.12
USER_GH=eyedeekay

version:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION) -d "I2P Router Checking CLI utility"

delete:
	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION)

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

build: fmt test clean
	cd ./i2cpcheck && go build $(GO_COMPILER_OPTS)
	cd ./i2cpcheck && GOOS=windows GOARCH=amd64 go build $(GO_COMPILER_OPTS) -buildmode=exe -o i2cpcheck.exe

test:
	go test -v

cli:
	./i2cpcheck/i2cpcheck && echo "Error condition confirmed"

clean:
	rm -f i2pccheck/i2cpcheck

fmt:
	find . -name '*.go' -exec gofmt -w -s {} \;

ZERO_VERSION=v1.16

i2p-zero:
	cd zerobundle && git clone https://github.com/i2p-zero/i2p-zero.git; \
		cd i2p-zero && \
		git fetch --all --tags && \
		git checkout $(ZERO_VERSION)

zero-build: i2p-zero
	cd zerobundle/i2p-zero && \
		./bin/build-all-and-zip.sh

zero-zip: 

zero-bundle: zero-zip
	cd zerobundle && \
		go run --tags generate ./gen/gen.go

ZEROB_VERSION=z9.46.12

zero-assets:
	gothub release -p -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "I2P Zero pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_windows.go" -f "zerobundle/windows/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_darwin.go" -f "zerobundle/mac/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_linux.go" -f "zerobundle/linux/assets.go"

I2PD_VERSION=2.30.0

i2pd-config:
	echo "using gcc : x86_64 : g++ ;" > ~/user-config.jam

BOOST_VERSION=1_72_0

boost: i2pdbundle/boost

i2pdbundle/boost:
	cd i2pdbundle && \
		wget -O boost.zip https://dl.bintray.com/boostorg/release/1.72.0/source/boost_$(BOOST_VERSION).zip && \
		unzip boost.zip && rm boost.zip && \
		mv boost_$(BOOST_VERSION) boost && \
		cd boost && \
			./bootstrap.sh && \
			./b2 toolset=gcc --prefix=$(PWD)/i2pdbundle/stage target-os=linux variant=release link=static runtime-link=static address-model=64 \
			--build-type=minimal --with-system --with-filesystem --with-program_options --with-date_time \
			--stagedir=$(PWD)/i2pdbundle/stage

#i2pdbundle/boost:
#	cd i2pdbundle && \
#		wget -O boost.zip https://dl.bintray.com/boostorg/release/1.72.0/source/boost_$(BOOST_VERSION).zip && \
#		unzip boost.zip && rm boost.zip && \
#		mv boost_$(BOOST_VERSION) boost && \
#		cd boost && \
#			./bootstrap.sh && \
#			./b2 toolset=gcc-mingw target-os=windows variant=release link=static runtime-link=static address-model=64 \
#			--build-type=minimal --with-system --with-filesystem --with-program_options --with-date_time \
#			--stagedir=stage-64

SSL_VERSION=1_0_2s

openssl-static:
	cd i2pdbundle && \
		git clone https://github.com/openssl/openssl; \
		cd openssl && \
		git checkout OpenSSL_$(SSL_VERSION) && \
		./Configure linux-x86_64 no-rc2 no-rc4 no-rc5 no-idea no-bf no-cast no-whirlpool no-md2 no-md4 no-ripemd no-mdc2 \
		  no-camellia no-seed no-comp no-krb5 no-gmp no-rfc3779 no-ec2m no-ssl2 no-jpake no-srp no-sctp no-srtp \
		  --prefix=$(PWD)/i2pdbundle/stage  && \
		make depend  && \
		make && \
		make install

zlib-static:
	cd i2pdbundle && \
		git clone https://github.com/madler/zlib; \
		cd zlib && \
		git checkout v1.2.8 && \
		CFLAGS=-O3 ./configure --static --64 --prefix=$(PWD)/i2pdbundle/stage && \
		make && \
		make install

MUPNP_VERSION=2.1

miniupnp-source: i2pdbundle/miniupnpc

i2pdbundle/miniupnpc:
	cd i2pdbundle && \
		wget -O miniupnp.tar.gz "http://miniupnp.free.fr/files/download.php?file=miniupnpc-$(MUPNP_VERSION).tar.gz" && \
		tar -xf miniupnp.tar.gz && rm miniupnp.tar.gz && \
		mv miniupnpc-$(MUPNP_VERSION) miniupnpc
	mkdir -p $(HOME)/dev
	make m

m:
	ln -sf $(PWD)/i2pdbundle/miniupnpc $(HOME)/dev/

hint-static:
	@echo set\(CMAKE_SYSTEM_NAME Windows\) | tee $(PWD)/i2pdbundle/toolchain-mingw.cmake
	@echo set\(CMAKE_C_COMPILER x86_64-w64-mingw32-gcc\) | tee -a $(PWD)/i2pdbundle/toolchain-mingw.cmake
	@echo set\(CMAKE_CXX_COMPILER x86_64-w64-mingw32\-g++\) | tee -a $(PWD)/i2pdbundle/toolchain-mingw.cmake
	@echo set\(CMAKE_RC_COMPILER x86_64-w64-mingw32-windres\) | tee -a $(PWD)/i2pdbundle/toolchain-mingw.cmake
	@echo set\(CMAKE_FIND_ROOT_PATH /usr/x86_64-w64-mingw32\) | tee -a $(PWD)/i2pdbundle/toolchain-mingw.cmake
	@echo set\(CMAKE_FIND_ROOT_PATH_MODE_PROGRAM NEVER\) | tee -a $(PWD)/i2pdbundle/toolchain-mingw.cmake

i2p-i2pd:
	cd i2pdbundle && git clone https://github.com/purplei2p/i2pd.git; \
		cd i2pd && \
		git fetch --all --tags && \
		git checkout $(I2PD_VERSION)

i2pd-build: i2p-i2pd
	cd i2pdbundle/i2pd && \
		make LIBDIR="$(WORK_DIR)boost/stage-64/lib $(WORK_DIR)stage/lib" \
			USE_STATIC=yes \
			LIB_ROOT=$(WORK_DIR)stage \
			BOOST_LIBRARYDIR=$(WORK_DIR)stage \
			OPENSSL_ROOT_DIR=$(WORK_DIR)stage \
			INCFLAGS="-I$(WORK_DIR)stage/include"

CMAKE_MAKE_PROGRAM=make
WORK_DIR=$(PWD)/i2pdbundle/

# -DWITH_UPNP=ON -DMINIUPNPC_LIBRARY=$(WORK_DIR)miniupnpc \
#export CC=x86_64-w64-mingw32-gcc
#export CXX=x86_64-w64-mingw32-g++

i2pd-build-2: i2p-i2pd
	cd i2pdbundle/i2pd && \
		CC=x86_64-w64-mingw32-gcc \
		CXX=x86_64-w64-mingw32-g++ \
		WITH_AESNI=ON \
		WITH_STATIC=ON \
		WITH_HARDENING=ON \
		ZLIB_ROOT=$(WORK_DIR)stage \
		BOOST_LIBRARYDIR:PATH=$(WORK_DIR)boost/stage-mingw-64/lib \
		OPENSSL_ROOT_DIR:PATH=$(WORK_DIR)stage \
		make

static-env: hint-static boost openssl-static zlib-static miniupnp-source

i2pd-build-windows: i2p-i2pd
	cd i2pdbundle/i2pd && \
		#rm -rf i2pd-mingw-64-build && \
		mkdir -p i2pd-mingw-64-build && \
		cd i2pd-mingw-64-build && \
			BOOST_ROOT=$(WORK_DIR)boost cmake -G 'Unix Makefiles' $(WORK_DIR)i2pd/build \
				-DBUILD_TYPE=Release \
				-DCMAKE_TOOLCHAIN_FILE=$(WORK_DIR)toolchain-mingw.cmake \
				-DWITH_AESNI=ON -DWITH_STATIC=ON -DWITH_HARDENING=ON \
				-DCMAKE_INSTALL_PREFIX:PATH=$(WORK_DIR)/i2pd/i2pd-mingw-64-static \
				-DZLIB_ROOT=$(WORK_DIR)stage \
				-DBOOST_LIBRARYDIR:PATH=$(WORK_DIR)boost/stage-mingw-64/lib \
				-DOPENSSL_ROOT_DIR:PATH=$(WORK_DIR)stage && \
	make && \
	x86_64-w64-mingw32-strip i2pd.exe

i2pd-zip:
	mkdir -p i2pdbundle/bundle
	wget -c -O i2pdbundle/bundle/osx.zip https://github.com/PurpleI2P/i2pd/releases/download/$(I2PD_VERSION)/i2pd_$(I2PD_VERSION)_osx.tar.gz
	cd i2pdbundle/bundle/ && tar xvf osx.zip
	wget -c -O i2pdbundle/bundle/win.zip https://github.com/PurpleI2P/i2pd/releases/download/$(I2PD_VERSION)/i2pd_$(I2PD_VERSION)_win64_mingw_avx_aesni.zip
	cd i2pdbundle/bundle/ && unzip win.zip

i2pd-bundle: i2pd-zip
	cd i2pdbundle #&& \
		#go run --tags generate ./gen/gen.go

I2PDB_VERSION=d9.46.12

i2pd-assets:
	#gothub release -p -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "i2pd C++ pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_windows.go" -f "i2pdbundle/windows/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_darwin.go" -f "i2pdbundle/mac/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_linux.go" -f "i2pdbundle/linux/assets.go"	