.PHONY: build view all

all: build-docker public
	@ ls -l main.wasm

build:
	tinygo build -o main.wasm -target wasm ./main.go

build-docker:
	# docker run --rm tinygo:confio ls /
	docker run --rm -v $(shell pwd):/code tinygo:confio tinygo build -o /code/main.wasm -target wasm /code/main.go


view:
	wasm-nm main.wasm

public:
	wasm-nm -e main.wasm
	wasm-nm -i main.wasm
