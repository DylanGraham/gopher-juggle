.PHONY: wasm

wasm: main.go
	GOOS=js GOARCH=wasm go build -o wasm/main.wasm -ldflags '-s -w' *.go

clean:
	rm -f wasm/main.wasm
