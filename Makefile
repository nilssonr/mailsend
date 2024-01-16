.PHONY: build_liunx
build_linux:
	GOOS=linux GOARCH=amd64 go build -o ./builds/linux/mailsend main.go