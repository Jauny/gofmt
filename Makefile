.PHONY: all	
all: wasm serve

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

.PHONY: serve
serve:
	bash -c "python3 -m http.server"
