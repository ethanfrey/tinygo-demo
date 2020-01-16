.PHONY: build view

build:
	tinygo build -o main.wasm -target wasm ./main.go

view:
	wasm-nm main.wasm

