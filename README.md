# Investigating tinygo

Try compiling with `make` and check out the imports/exports of the file.

You will need to [install tinygo](https://tinygo.org/getting-started/) first.
For Ubuntu, this is:

```
wget https://github.com/tinygo-org/tinygo/releases/download/v0.11.0/tinygo_0.11.0_amd64.deb
sudo dpkg -i tinygo_0.11.0_amd64.deb
```

## Research

This is very close to what we want for generating wasmer-compatible wasm: https://github.com/tinygo-org/tinygo/issues/207

These are all the directives we can use on functions: https://github.com/tinygo-org/tinygo/blob/master/ir/ir.go#L28-L38

## Building stuff

Check out and build and docker file:

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

You see the exports we wrote:

```
e add
e bcounter
e counter
e think
e relic
```

And no auto-generated imports. Now, let's try this with wasmer

**TODO** automate this build system a bit better... at least a tagged docker image I publish would be good

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
