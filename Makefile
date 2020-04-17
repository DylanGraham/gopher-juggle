.PHONY: wasm

wasm: main.go
	GOOS=js GOARCH=wasm go build -o wasm/main.wasm -ldflags '-s -w' main.go

clean:
	rm wasm/main.wasm
