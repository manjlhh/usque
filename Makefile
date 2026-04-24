clean:
	rm /tmp/usque
	rm /tmp/uq

build:
	go build -o /tmp/usque -ldflags="-s -w" .
	upx --best --lzma -o /tmp/uq /tmp/usque
