
GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

build: fmt test clean
	cd ./i2cpcheck && go build $(GO_COMPILER_OPTS) && \
		GOOS=windows GOARCH=amd64 go build $(GO_COMPILER_OPTS) -buildmode=exe -o i2cpcheck.exe

test:
	go test -v

clean:
	rm -f i2pccheck/i2cpcheck

fmt:
	find . -name '*.go' -exec gofmt -w -s {} \;
