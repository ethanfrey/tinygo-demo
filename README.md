# Investigating tinygo

Some experiments with what it can do

## Research

This is very close to what we want for generating wasmer-compatible wasm: https://github.com/tinygo-org/tinygo/issues/207

These are all the directives we can use on functions: https://github.com/tinygo-org/tinygo/blob/master/ir/ir.go#L28-L38

## Prepare tooling

You can use the standard file:

```
docker pull tinygo/tinygo:0.13.1
```

You will also want [`wasm-nm`](https://crates.io/crates/wasm-nm) to investigate the wasm output.

```
cargo install wasm-nm
```

## Building stuff

Now, let's build this example and check it out.
Go to this directory and:

```sh
make
```

You see the exports we wrote (properly renamed):

```
wasm-nm -e main.wasm
e __wasm_call_ctors
e memset
e _start
e go_scheduler
e resume
e init
e query
e migrate
e /code/main.go.main
```

And imports:

```
wasm-nm -i main.wasm
i fd_write
```

You can dig in more with `make view`.

## Custom compiles

If you want to hack the exports, check out and build and docker file:

```sh
git clone https://github.com/confio/tinygo
cd tinygo
git checkout cosmwasm-v2
docker build . -f Dockerfile.wasm -t cosmwasm/tinygo:latest
```

Now, try compiling and testing with:

```
make build-cosmwasm
make view
```

Sample Result (note fewer exports and importantly no imports required):

```
e __wasm_call_ctors
e memset
e _start
e init
e query
e migrate
e /code/main.go.main
e handle
-rwxr-xr-x 1 root root 5335 jun 23 21:17 main.wasm
```

And for even smaller (production) builds, try:

```
make build-cosmwasm-trim
make view
```

Sample Result (look at the code size just shrinking...):

``` 
e __wasm_call_ctors
e memset
e _start
e init
e query
e migrate
e /code/main.go.main
e handle
-rwxr-xr-x 1 root root 1022 jun 23 21:17 main.wasm
```

## JSON

I tried to add json parsing with "encoding/json" and the compiler blew up.
Reflect is not supported.

This is an open issue: https://github.com/tinygo-org/tinygo/issues/93#issuecomment-552050452

The only "workaround" I know of is to use https://github.com/vugu/vjson which does compile with TinyGo

Tracking cgo, which we need some to be able to parse structs (like Buffer): https://github.com/tinygo-org/tinygo/issues/60
It seems there is progess, but it may still require the `-no-debug` flag, and if you have any trouble, look at those issues for hints.
