package renderer

import (
	"fmt"
	"strconv"
)

type canvas struct {
	width  int
	height int
	cell   [][]string
}

func NewCanvas(width int, height int) *canvas {
	cells := make([][]string, height)
	for i := range cells {
		cells[i] = make([]string, width)
	}

	return &canvas{
		width:  width,
		height: height,
		cell:   cells,
	}
}

func (c *canvas) Draw() *canvas {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			c.cell[i][j] = " "
		}
	}

	return c
}

func (c *canvas) DrawChart(numbers []int) {
	temp := make([]int, len(numbers))
	copy(temp, numbers)

	for i := c.height - 1; i >= 0; i-- {
		k := 0
		for j := 0; j < c.width; j++ {
			if j%2 == 0 && i == c.height-1 {
				c.cell[i][j] = strconv.Itoa(temp[k])
			} else if j%2 == 0 && temp[k] != 0 {
				c.cell[i][j] = "|"
				temp[k] -= 1
			} else {
				c.cell[i][j] = " "
			}

			if j%2 == 0 {
				k++
			}
		}
	}
}

func (c *canvas) Display() {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			fmt.Printf("%+v", c.cell[i][j])
		}
		println()
	}
}
