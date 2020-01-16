.PHONY: build view all

all: build public

build:
	tinygo build -o main.wasm -target wasm ./main.go

view:
	wasm-nm main.wasm

public:
	wasm-nm -e main.wasm
	wasm-nm -i main.wasm
