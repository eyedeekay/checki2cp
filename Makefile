

GO111MODULE=on

VERSION=0.0.12
USER_GH=eyedeekay

version:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION) -d "I2P Router Checking CLI utility"

delete:
	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION)

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

btest: fmt
	cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && go build $(GO_COMPILER_OPTS) && ./go-i2pd

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

ZEROB_VERSION=z9.45.12

zero-assets:
	gothub release -p -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "I2P Zero pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_windows.go" -f "zerobundle/windows/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_darwin.go" -f "zerobundle/mac/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(ZEROB_VERSION) -n "assets_linux.go" -f "zerobundle/linux/assets.go"

I2PD_VERSION=2.31.0

i2pd-clean:
	rm -rf i2pdbundle/appimage i2pdbundle/osx i2pdbundle/win

i2pd-zip: i2pd-clean i2pd-linux
	mkdir -p i2pdbundle/mac i2pdbundle/win i2pdbundle/test i2pdbundle/test/subtest i2pdbundle/test/subtest/subsubtest i2pdbundle/test/subsubsubtest
	wget -c -qO i2pdbundle/mac.zip https://github.com/PurpleI2P/i2pd/releases/download/$(I2PD_VERSION)/i2pd_$(I2PD_VERSION)_osx.tar.gz
	cd i2pdbundle/mac && tar xvf ../mac.zip
	wget -c -qO i2pdbundle/win.zip https://github.com/PurpleI2P/i2pd/releases/download/$(I2PD_VERSION)/i2pd_$(I2PD_VERSION)_win64_mingw_avx_aesni.zip
	cd i2pdbundle/win && unzip ../win.zip
	touch i2pdbundle/test/test.txt \
		i2pdbundle/test/subtest/test.txt \
		i2pdbundle/test/subtest/subsubtest/test.txt \
		i2pdbundle/test/subsubsubtest/test.txt \
		i2pdbundle/test/test_other.txt \
		i2pdbundle/test/subtest/test_other.txt \
		i2pdbundle/test/subtest/subsubtest/test_other.txt \
		i2pdbundle/test/subsubsubtest/test_other.txt

i2pd-linux:
	mkdir -p i2pdbundle/linux/lib
	cd $(WORK_DIR)/i2pd-static-64-build/ && tar czvf ../../../i2pdbundle/linux/i2pd.tar.gz ./i2pd
	cp /lib64/ld-linux-x86-64.so.2 i2pdbundle/linux/lib
	cp /lib/x86_64-linux-gnu/libc.so.6 i2pdbundle/linux/lib
	cp /lib/x86_64-linux-gnu/libdl.so.2 i2pdbundle/linux/lib
	cp /lib/x86_64-linux-gnu/libgcc_s.so.1 i2pdbundle/linux/lib
	cp /lib/x86_64-linux-gnu/libm.so.6 i2pdbundle/linux/lib
	cp /lib/x86_64-linux-gnu/libpthread.so.0 i2pdbundle/linux/lib
	cp /usr/lib/x86_64-linux-gnu/libstdc++.so.6 i2pdbundle/linux/lib
	cd i2pdbundle/linux/lib && tar czvf ../lib.tar.gz .
	rm -rf i2pdbundle/linux/lib


i2p-i2pd:
	cd i2pdbundle && git clone https://github.com/purplei2p/i2pd.git; \
		cd i2pd && \
		git fetch --all --tags && \
		git checkout $(I2PD_VERSION)


i2pd-build-static: i2p-i2pd
	cd i2pdbundle/i2pd/ && make clean && \
		USE_STATIC=yes make USE_STATIC=yes


WORK_DIR=$(PWD)/i2pdbundle/i2pd

i2pd-build: i2p-i2pd
	cd $(WORK_DIR) && \
		rm -rf $(WORK_DIR)/i2pd-static-64-build && \
		mkdir -p $(WORK_DIR)/i2pd-static-64-build && \
		cd $(WORK_DIR)/i2pd-static-64-build && \
			cmake -G 'Unix Makefiles' $(WORK_DIR)/build \
				-DBUILD_TYPE=Release \
				-DWITH_STATIC=ON -DWITH_HARDENING=ON -DWITH_UPNP=ON \
				-DCMAKE_INSTALL_PREFIX:PATH=$(WORK_DIR)/i2pd-static-64-build && \
	make CXXFLAGS=-static

i2pd-bundle: i2pd-zip
	cd i2pdbundle && \
		go run --tags generate ./gen.go

I2PDB_VERSION=d9.45.12

i2pd-assets:
	#gothub release -p -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "i2pd C++ pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_windows.go" -f "i2pdbundle/windows/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_darwin.go" -f "i2pdbundle/mac/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_linux.go" -f "i2pdbundle/linux/assets.go"	
