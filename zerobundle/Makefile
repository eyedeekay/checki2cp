
gen: get write test

get:
	wget -c -O i2p-zero-linux.v1.17.zip https://github.com/i2p-zero/i2p-zero/releases/download/v1.17/i2p-zero-linux.v1.17.zip
	wget -c -O i2p-zero-win.v1.17.zip https://github.com/i2p-zero/i2p-zero/releases/download/v1.17/i2p-zero-win.v1.17.zip
	wget -c -O i2p-zero-darwin.v1.17.zip https://github.com/i2p-zero/i2p-zero/releases/download/v1.17/i2p-zero-mac.v1.17.zip

write:
	go gen --tags=generate gen.go

test:
	cd import/ && go test
