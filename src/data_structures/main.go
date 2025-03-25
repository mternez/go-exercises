package main

import (
	"data_structures/rope"
	"fmt"
)

func main() {
	root := rope.NewRope("Hello World!")
	fmt.Println(rope.IterativeString(root))
}
