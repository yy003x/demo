package main

import (
	"fmt"
)

func escapeToHeap() *int {
	b := 42
	return &b // b 逃逸到堆上
}

func stayOnStack() {
	c := 100 // c 在栈上分配
	fmt.Println(c)
}

func main() {
	escapeToHeap()
	stayOnStack()
}
