package main

func main() {
}

//go:export add
func add(a, b int) int {
	return a + b
}

// why is this not exported???
//go:export relic
func relic(a int) int {
	l := counter("foobar")
	return 10 * a + l
}

// Name mismatch cause no wasm export
//go:export think
func foobar(a, b int) int {
	return a + b
}

// []byte cause no wasm export
//go:export bcounter
func bcounter(a []byte) int {
	return len(a)
}

//go:export counter
func counter(a string) int {
	return len(a)
}
