wasm: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm -ldflags '-s -w' main.go
	mv main.wasm wasm
