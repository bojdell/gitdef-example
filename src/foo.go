package main

import (
	"fmt"
)

func foo(x int) int {
	return bar(x) + 1
}

func main() {
	fmt.Println("foobar", foo(1))
}
