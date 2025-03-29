package main

import (
	"data_structures/rope"
)

func main() {
	a := rope.NewRope("Ropes are very cool you should try to use them", 1)
	rope.Rebalance(a)
}
