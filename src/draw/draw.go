package draw

import (
	"fmt"
	"life/src/sim"
)

func Split(f *sim.Frame) []string {

	lines := make([]string, f.Height)

	for y := range lines {
		line := ""
		for x := range f.Cells[y] {
			if f.Cells[x][y] {
				line += "\033[32m * \033[0m"
			} else {
				line += "   "
			}
		}
		lines[y] = "|" + line + "|"
	}

	separator := " "

	for i := 0; i < f.Width; i++ {
		separator += "---"
	}

	lines[0], lines[f.Height-1] = separator, separator

	return lines
}

func Output(f *sim.Frame) {
	lines := Split(f)
	fmt.Printf("\033[%dA", f.Height+4)
	fmt.Println("\033[0K")

	for line := range lines {
		fmt.Println(lines[line])
	}

}
