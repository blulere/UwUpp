all: windows linux macos silicon

windows: uwupp-go/main.go
	echo Building a Windows executable...
	GOOS=windows go build -o uwupp-go/bin/uwupp-windows.exe uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
	echo Finished!
linux: uwupp-go/main.go
	echo Building a Linux executable...
	echo Note that you may have to enable executing this file as a program in permissions settings.
	GOOS=linux go build -o uwupp-go/bin/uwupp-linux uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
	echo Finished!
macos: uwupp-go/main.go
	echo Building a macOS executable...
	GOOS=darwin GOARCH=amd64 go build -o uwupp-go/bin/uwupp-macos uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
	echo Finished!
silicon: uwupp-go/main.go
	echo Building a macOS Silicon executable...
	GOOS=darwin GOARCH=arm64 go build -o uwupp-go/bin/uwupp-sillicon uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
	echo Finished!

clean:
	rm -rf uwupp-go/bin
	del uwupp-go/bin