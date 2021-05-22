

GO111MODULE=on

VERSION=0.0.18
USER_GH=eyedeekay

version:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION) -d "I2P Router Checking CLI utility"

delete:
	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION)

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

launchers: fmt test clean
	cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && go build $(GO_COMPILER_OPTS)
	cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && GOOS=windows go build -o go-i2pd.exe $(GO_COMPILER_OPTS)
	cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && GOOS=windows go build -o go-i2pd-32.exe $(GO_COMPILER_OPTS)
	#cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && GOOS=darwin go build $(GO_COMPILER_OPTS)


btest: fmt
	cd ./go-i2pd && rm -rf i2pd lib i2pd.pid && go build $(GO_COMPILER_OPTS) && ./go-i2pd

build: fmt test clean
	cd ./i2cpcheck && go build $(GO_COMPILER_OPTS)
	cd ./i2cpcheck && GOOS=windows GOARCH=amd64 go build $(GO_COMPILER_OPTS) -buildmode=exe -o i2cpcheck.exe

test:
	go test -v
	go test -v ./controlcheck
	go test -v ./samcheck
	go test -v ./proxycheck

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

	#gothub release -p -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "i2pd C++ pre-encoded assets" -d "assets.go file containing a zipped bundle of I2P Zero"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_windows.go" -f "i2pdbundle/windows/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_darwin.go" -f "i2pdbundle/mac/assets.go"
	#gothub upload -R -u eyedeekay -r "checki2cp" -t $(I2PD_VERSION) -n "assets_linux.go" -f "i2pdbundle/linux_amd64/assets.go"	
