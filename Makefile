help:
	@echo "Tunalang Makefile (Requires Go v1.16+)"
	@echo "'help' - Displays the command usage"
	@echo "'test'"
	@echo "'build' - Builds the Tunalang binary" 
	@echo "'install' - Installs Tunalang"
	@echo "'build-windows' - Builds the Tunalang binary for Windows platforms"
	@echo "'build-darwin' - Builds the Tunalang binary for MacOSX platforms"
	@echo "'build-linux' - Builds the Tunalang binary for Linux platforms"
	@echo "'build-all' - Builds the Tunalang binary for all platforms"

test:
	@echo Testing Tunalang Components
	@go test -v ./...
	@echo Test Complete

build:
	@echo Building Tunalang binary
	@go build .
	@echo Build Complete. Run './tunalang(.exe)'

install:
	@echo Installing Tunalang
	@go install .
	@echo install Complete. Run 'tunalang'.

build-windows:
	@echo Cross Compiling for Windows x86
	@GOOS=windows GOARCH=386 go build -o ./bin/tunalang-windows-x32.exe

	@echo Cross Compiling for Windows x64
	@GOOS=windows GOARCH=amd64 go build -o ./bin/tunalang-windows-x64.exe

	@echo Cross Compiled for Windows platforms

build-macosx:
	@echo Cross Compiling for MacOSX x64
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/tunalang-macosx-x64

	@echo Cross Compiled for MacOSX platforms

build-linux:
	@echo Cross Compiling for Linux x32
	@GOOS=linux GOARCH=386 go build -o ./bin/tunalang-linux-x32

	@echo Cross Compiling for Linux x64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/tunalang-linux-x64

	@echo Cross Compiled for Linux x86 platforms

build-linux-arm:
	@echo Cross Compiling for Linux Arm32
	@GOOS=linux GOARCH=arm go build -o ./bin/tunalang-linux-arm32

	@echo Cross Compiling for Linux Arm64
	@GOOS=linux GOARCH=arm64 go build -o ./bin/tunalang-linux-arm64

	@echo Cross Compiled for Linux ARM platforms

build-all: build-windows build-macosx build-linux build-linux-arm
	@echo Cross Compiled for all platforms