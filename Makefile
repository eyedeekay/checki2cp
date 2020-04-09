

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
	echo "using gcc : mingw : x86_64-w64-mingw32-g++ ;" > ~/user-config.jam

BOOST_VERSION=1_72_0

boost-windows:
	cd i2pdbundle && \
		wget -O boost.zip https://dl.bintray.com/boostorg/release/1.72.0/source/boost_$BOOST_VERSION.zip && \
		unzip boost.zip && \
		mv boost_$(BOOST_VERSION) boost && \
		cd boost #&& \
			#./bootstrap.sh &&
			#./b2 toolset=gcc-mingw target-os=windows variant=release link=static runtime-link=static address-model=64 \
			#--build-type=minimal --with-filesystem --with-program_options --with-date_time \
			#--stagedir=stage-mingw-64

openssl-windows:
	cd i2pdbundle && \
		git clone https://github.com/openssl/openssl; \
		cd openssl && \
		git checkout OpenSSL_1_0_2s && \
		./Configure mingw64 no-rc2 no-rc4 no-rc5 no-idea no-bf no-cast no-whirlpool no-md2 no-md4 no-ripemd no-mdc2 \
		  no-camellia no-seed no-comp no-krb5 no-gmp no-rfc3779 no-ec2m no-ssl2 no-jpake no-srp no-sctp no-srtp \
		  --prefix=~/dev/stage --cross-compile-prefix=x86_64-w64-mingw32-  && \
		make depend  && \
		make  && \
		make install

zlib-windows:
	cd i2pdbundle && \
		git clone https://github.com/madler/zlib; \
		cd zlib && \
		git checkout v1.2.8 && \
		CC=x86_64-w64-mingw32-gcc CFLAGS=-O3 ./configure --static --64 --prefix=~/dev/stage && \
		make && \
		make install

i2p-i2pd:
	cd i2pdbundle && git clone https://github.com/purplei2p/i2pd.git; \
		cd i2pd && \
		git fetch --all --tags && \
		git checkout $(I2PD_VERSION)

i2pd-build: i2p-zero
	cd i2pdbundle/i2pd && \
		make

i2pd-zip: 

i2pd-bundle: i2pd-zip
	cd i2pdbundle #&& \
		#go run --tags generate ./gen/gen.go

I2PDB_VERSION=d9.46.12

i2pd-assets:
	#gothub release -p -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "i2pd C++ pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_windows.go" -f "i2pdbundle/windows/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_darwin.go" -f "i2pdbundle/mac/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_linux.go" -f "i2pdbundle/linux/assets.go"	