package main

import (
	"data_structures/rope"
	"fmt"
)

func main() {
	root := rope.NewRope("Ropes are very cool you should try to use them", 1)
	for _, leaf := range rope.Rebalance(root.CollectLeaves()).CollectLeaves() {
		fmt.Println(leaf)
	}

}
