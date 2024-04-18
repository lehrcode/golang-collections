package main

import (
	"fmt"
	"github.com/lehrcode/golang-collections"
	"github.com/lehrcode/golang-collections/linkedlist"
	"github.com/lehrcode/golang-collections/slicelist"
)

func main() {
	var stack collections.Stack[int] = slicelist.New[int]()
	stack.Push(5)
	stack.Push(123)
	stack.Push(3)
	stack.Pop()
	fmt.Println(stack)

	var stack2 collections.Stack[int] = linkedlist.New[int]()
	stack2.Push(5)
	stack2.Push(7)
	stack2.Push(9)
	stack2.Push(11)
	stack2.Pop()
	stack2.Push(27)
	fmt.Println(stack2)
}
