package main

import (
	"data_structures/rope"
	"fmt"
)

func main() {
	root := rope.NewRope("Ropes are very cool you try to use them")
	for _, leaf := range root.CollectLeaves() {
		fmt.Println(leaf.Data())
	}
	fmt.Println(root.Weight())
}
