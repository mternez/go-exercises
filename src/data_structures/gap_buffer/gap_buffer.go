package gapbuffer

import "fmt"

const empty rune = 0

type GapBuffer struct {
	buffer     []rune
	bufferSize int
	gapSize    int
	start      int
	end        int
}

func NewGapBuffer(size int, gap int) *GapBuffer {
	return &GapBuffer{
		buffer:     make([]rune, size),
		bufferSize: size,
		gapSize:    gap,
		start:      0,
		end:        gap - 1,
	}
}

func (b *GapBuffer) Insert(c rune, pos int) {

	b.moveGapTo(pos)
	b.buffer[pos] = c
	b.start++
	if b.start == b.end {
		b.reallocate()
	}
}

func (b *GapBuffer) Delete(pos int) {
	b.moveGapTo(pos)
	b.buffer[pos] = empty
	if b.start > 0 {
		b.start--
	}
}

func (b *GapBuffer) moveGapTo(pos int) {
	if pos < 0 || pos > b.bufferSize {
		fmt.Println("moveGapTo: out of bounds")
		return
	}

	currentGapSize := (b.end - b.start)
	if pos < b.start {
		// Move left
		for i := b.start - 1; i >= pos && i >= 0; i-- {
			b.buffer[i+currentGapSize] = b.buffer[i]
			b.buffer[i] = empty
		}
	} else if pos > b.start {
		// Move right
		for i := b.start; i < pos && i >= b.bufferSize; i++ {
			b.buffer[i] = b.buffer[i+currentGapSize]
			b.buffer[i+currentGapSize] = empty
		}
	}
	b.start = pos
	b.end = b.start + currentGapSize
}

func (b *GapBuffer) reallocate() {
	fmt.Println("INFO:reallocating")
	newBuffer := make([]rune, b.bufferSize+b.gapSize)
	copy(newBuffer[:b.start], b.buffer[:b.start])
	copy(newBuffer[b.start+b.gapSize:], b.buffer[b.end:])
	b.buffer = newBuffer
	b.bufferSize += b.gapSize
	b.end = b.start + b.gapSize
}

func (b *GapBuffer) Buffer() []rune {
	return b.buffer
}

func (b *GapBuffer) String() string {
	return string(b.buffer)
}

func (b *GapBuffer) DrawGap() {
	for ind := b.start; ind < b.end; ind++ {
		b.buffer[ind] = '_'
	}
	b.buffer[b.start] = '['
	b.buffer[b.end] = ']'
}

func (b *GapBuffer) EraseDrawnGap() {
	for ind := b.start; ind < b.end; ind++ {
		b.buffer[ind] = 0
	}
	b.buffer[b.start] = 0
	b.buffer[b.end] = 0
}

func (b *GapBuffer) PrintWithVisibleGap(header string) {
	fmt.Printf("===============\n%s\n===============\nstart:%d,end:%d\n", header, b.start, b.end)
	b.DrawGap()
	for _, r := range b.buffer {
		if r == '_' {
			fmt.Printf(" %s ", "|")
		} else if r == empty {
			fmt.Printf(" %s ", ".")
		} else {
			fmt.Printf(" %s ", string(r))
		}
	}
	fmt.Println("")
	b.EraseDrawnGap()
	fmt.Printf("===============\n")
}
