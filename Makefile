.PHONY: build

build:
	tinygo build -o main.wasm -target wasm ./main.go

