package main

import (
	"encoding/json"
)

func main() {
}

//go:export add
func add(a, b int) int {
	return a + b
}

//go:export relic
func relic(a int) int {
	l := counter("foobar")
	return 10 * a + l
}

type model struct {
	Age int32 `json:"age"`
}

//go:export think
func foobar() int {
	var data model
	input := `{"age":17}`
	err := json.Unmarshal(input, &data)
	if err != nil {
		return -1
	}
	return data.Age
}

//go:export bcounter
func bcounter(a []byte) int {
	return len(a)
}

//go:export counter
func counter(a string) int {
	return len(a)
}
