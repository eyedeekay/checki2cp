

GO111MODULE=on

VERSION=0.33.8
USER_GH=eyedeekay

version:
	github-release release -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION) -d "I2P Router Checking CLI utility and libraries"

delete:
	github-release delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r checki2cp -t v$(VERSION)

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

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
	find . -name '*.go' -exec gofumpt -extra -w -s {} \;
