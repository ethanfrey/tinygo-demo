package main

func main() {
}

//go:export init
func add(a, b int) int {
	return a + b
}

//go:wasm-module env
//go:export handle
func relic(a int) int {
	l := counter("foobar")
	return 10 * a + l
}

//go:wasm-module foo
//go:export query
func bcounter(a []byte) int {
	return len(a)
}

//go:export migrate
func counter(a string) int {
	return len(a)
}
