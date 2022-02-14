package sim

import (
	"math/rand"
)

type Frame struct {
	Cells  [][]bool
	Width  int
	Height int
}

func NewFrame(width int, height int) *Frame {
	c := make([][]bool, width)
	for i := range c {
		c[i] = make([]bool, height)
	}
	return &Frame{Cells: c, Width: width, Height: height}
}

// Populate array with random values, determined by `density`.
func Populate(f *Frame, density float64) {
	for x := range f.Cells {
		for y := range f.Cells[x] {
			f.Cells[x][y] = rand.Float64() < density
		}
	}
}

func NextFrame(f Frame) *Frame {
	next := NewFrame(f.Width, f.Height)

	updateCell := func(x int, y int) bool {
		// Find # of neighbors
		neighbors := 0
		// 3 rows to check for neighbors
		for i := -1; i <= 1; i++ {
			// Iterate through row
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if x+i < 0 || x+i >= f.Width {
					continue
				}
				if y+j < 0 || y+j >= f.Height {
					continue
				}
				if f.Cells[x+i][y+j] {
					neighbors++
				}
			}
		}

		var newState bool

		if (neighbors == 2 || neighbors == 3) && f.Cells[x][y] {
			newState = true
		} else if neighbors == 3 && !f.Cells[x][y] {
			newState = true
		} else {
			newState = false
		}

		return newState

	}

	for x := range f.Cells {
		for y := range f.Cells[x] {
			next.Cells[x][y] = updateCell(x, y)
		}
	}

	// fmt.Println(next)
	// fmt.Println(f)

	return next
}
