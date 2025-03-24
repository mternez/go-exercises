package main

import (
	"deck/deck"
	"fmt"
)

type MyFamily struct {
	deck.Family
}

type MyCard struct {
	deck.Card
}

func main() {
	c := &MyCard{}
	c.SetName("ACE")
	c.SetValue(244)
	f := &MyFamily{}
	f.SetName("HEART")
	c.SetFamily(f)
	fmt.Print(c)
}
