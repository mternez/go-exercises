package main

import (
	gapbuffer "data_structures/gap_buffer"
	"fmt"
)

func main() {
	a := gapbuffer.NewGapBuffer(50, 10)
	a.Insert('1', 10)
	a.PrintWithVisibleGap("After Insert '1' at 10")
	a.Insert('2', 11)
	a.PrintWithVisibleGap("After Insert '2' at 11")
	a.Insert('3', 12)
	a.PrintWithVisibleGap("After Insert '3' at 12")
	a.Insert('6', 30)
	a.PrintWithVisibleGap("After Insert '6' at 30")
	a.Insert('7', 31)
	a.PrintWithVisibleGap("After Insert '7' at 31")
	a.Insert('8', 32)
	a.PrintWithVisibleGap("After Insert '8' at 32")
	a.Insert('9', 33)
	a.PrintWithVisibleGap("After Insert '9' at 33")
	a.Insert('A', 34)
	a.PrintWithVisibleGap("After Insert 'A' at 34")
	a.Insert('B', 35)
	a.PrintWithVisibleGap("After Insert 'B' at 35")
	a.Insert('C', 36)
	a.PrintWithVisibleGap("After Insert 'C' at 36")
	a.Insert('D', 36)
	a.PrintWithVisibleGap("After Insert 'D' at 36")
	a.Insert('E', 36)
	a.PrintWithVisibleGap("After Insert 'D' at 36")
	fmt.Println(string(a.Buffer()))
}
