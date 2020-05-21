MAIN_PATH = cmd/go-hex/main.go

build:
	go build -o bin/main $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

compile-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 $(MAIN_PATH)

compile-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 $(MAIN_PATH)

compile-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 $(MAIN_PATH)