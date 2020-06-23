# Investigating tinygo

Some experiments with what it can do

## Research

This is very close to what we want for generating wasmer-compatible wasm: https://github.com/tinygo-org/tinygo/issues/207

These are all the directives we can use on functions: https://github.com/tinygo-org/tinygo/blob/master/ir/ir.go#L28-L38

## Building stuff

You can use the standard file:

```
docker pull tinygo/tinygo:0.13.1
```

Or if you want to hack the exports (runtime...), check out and build and docker file:
(this is still on 0.11, so needs an update for sure)

```sh
git clone https://github.com/confio/tinygo
cd tinygo
git checkout cosmwasm
docker build . -t tinygo:confio
```

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

Note that I have a [demo commit to remove some of this](https://github.com/confio/tinygo/commit/95a4707853f8c8522f6a51db14c8306727a3776f),
but need to figure out the proper way to bring this upstream.

## JSON

I tried to add json parsing with "encoding/json" and the compiler blew up.
Reflect is not supported.

This is an open issue: https://github.com/tinygo-org/tinygo/issues/93#issuecomment-552050452

I don't know if there are any workarounds.

Tracking cgo, which we need some to be able to parse structs (like Buffer): https://github.com/tinygo-org/tinygo/issues/60

## Wasmer

* TODO: Write a simple tester in go-ext-wasm and see if we can run functions
* TODO: Try to implement standard cosmwasm api
* TODO: Connect with rust cosmwasm
* TODO: Implement compatibility

At the end, we should be able to get a go hello world working in cosmwasm

Let's check feasibility, then pitch
