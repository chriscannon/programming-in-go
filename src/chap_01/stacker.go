package main

import (
	"chap_01/stack"
	"fmt"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hey")
	haystack.Push("ho")
	haystack.Push("let's go")

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
