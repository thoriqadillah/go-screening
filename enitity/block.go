package enitity

type block struct {
	X    int
	Y    int
	Char string
}

func NewBlock(x int, y int, char string) block {
	return block{x, y, char}
}
