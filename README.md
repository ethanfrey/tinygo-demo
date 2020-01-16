# Investigating tinygo

Try compiling with `make` and check out the imports/exports of the file.

You will need to [install tinygo](https://tinygo.org/getting-started/) first.

## Research

This is very close to what we want for generating wasmer-compatible wasm: https://github.com/tinygo-org/tinygo/issues/207

These are all the directives we can use on functions: https://github.com/tinygo-org/tinygo/blob/master/ir/ir.go#L28-L38
