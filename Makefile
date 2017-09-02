
.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/confirm_darwin_amd_64 .
	GOOS=linux GOARCH=amd64 go build -v -o bin/confirm_linux_amd_64 .
