windows:
	GOOS=windows go build -o uwupp-go/bin/uwupp-windows.exe uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
linux:
	GOOS=linux go build -o uwupp-go/bin/uwupp-linux uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
macos:
	GOOS=darwin GOARCH=amd64 go build -o uwupp-go/bin/uwupp-macos uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
sillicon:
	GOOS=darwin GOARCH=arm64 go build -o uwupp-go/bin/uwupp-sillicon uwupp-go/main.go uwupp-go/interpreter.go uwupp-go/exits.go
