package main

import (
	"fmt"
	"github.com/lehrcode/golang-collections"
	"github.com/lehrcode/golang-collections/linkedlist"
)

func main() {
	var list collections.List[string] = linkedlist.New[string]()

	fmt.Println(list)
}
