package enitity

import (
	"fmt"
	"time"
)

type cursor struct{}

func NewCursor() cursor {
	return cursor{}
}

func (c *cursor) refresh() {
	fmt.Print("\u001b[H")
}

func (c *cursor) Render(canvas *canvas) {
	time.Sleep(time.Second / 60)
	canvas.Display()
	c.refresh()
}
