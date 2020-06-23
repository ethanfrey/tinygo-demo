.PHONY: build view all

DOCKER_IMAGE=tinygo/tinygo:0.13.1

all: build-docker public
	@ ls -l main.wasm

build:
	tinygo build -o main.wasm -target wasm ./main.go

build-docker:
	# docker run --rm tinygo:confio ls /
	docker run --rm -v $(shell pwd):/code $(DOCKER_IMAGE) tinygo build -o /code/main.wasm -target wasm /code/main.go

view:
	@ wasm-nm main.wasm
	@ ls -l main.wasm

public:
	wasm-nm -e main.wasm
	wasm-nm -i main.wasm
