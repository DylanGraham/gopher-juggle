.PHONY: wasm

wasm: main.go
	GOOS=js GOARCH=wasm go build -o wasm/main.wasm -ldflags '-s -w' main.go
	cp *.wav *.png wasm/

clean:
	rm -f wasm/main.wasm wasm/*.wav wasm/*.png
