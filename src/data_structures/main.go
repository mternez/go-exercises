package main

import (
	gapbuffer "data_structures/gap_buffer"
	"fmt"
)

func main() {
	a := gapbuffer.NewGapBuffer(50, 10)
	str := "Gap buffers are very cool, you should try them"
	a.MoveCursor(6)
	a.PrintWithVisibleGap("After moving cursor to position 6")
	for _, r := range str {
		a.Insert(r)
		a.PrintWithVisibleGap("After inserting '" + string(r) + "'")
	}
	fmt.Println(string(a.Buffer()))
}
