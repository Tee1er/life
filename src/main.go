package main

import (
	"life/src/draw"
	"life/src/sim"
	"time"
)

func main() {
	// Create a new frame to store the values
	f := sim.NewFrame(32, 32)
	sim.Populate(f, 0.5)

	// output(draw.Split(f))

	for {
		f = sim.NextFrame(*f)
		time.Sleep(500 * time.Millisecond)
		draw.Output(f)
	}
	// Update frames
}
