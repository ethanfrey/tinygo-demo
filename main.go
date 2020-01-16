package main

func main() {
}

//go:export add
func add(a, b int) int {
	return a + b
}

//go:export multiply
func multiply(a, b int) int {
	return a * b
}

