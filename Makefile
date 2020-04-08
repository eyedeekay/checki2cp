

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

LAUNCH_VERSION=9.46.12

zero-assets:
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(LAUNCH_VERSION) -n "windows-assets.go" -f "zerobundle/windows/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(LAUNCH_VERSION) -n "mac-assets.go" -f "zerobundle/mac/assets.go"
	gothub upload -R -u eyedeekay -r "checki2cp" -t $(LAUNCH_VERSION) -n "linux-assets.go" -f "zerobundle/linux/assets.go"