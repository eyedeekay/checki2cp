
build: test clean
	cd ./i2cpcheck && go build

test:
	go test -v

clean:
	rm -f i2pccheck/i2cpcheck
