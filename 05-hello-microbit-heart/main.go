package main

import (
	"image/color"

	"tinygo.org/x/drivers/microbitmatrix"
)

func main() {
	display := microbitmatrix.New()
	display.Configure(
		microbitmatrix.Config{},
	)
	display.DisableAll()
	display.ClearDisplay()

	heart := [5][5]int8{
		{0, 1, 0, 1, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
	}

	c := color.RGBA{255, 255, 255, 0}

	for row := int16(0); row < 5; row++ {
		for col := int16(0); col < 5; col++ {
			if heart[row][col] == 1 {
				display.SetPixel(col, row, c)
			}
		}
	}

	for {
		// show heart forever
		display.Display()
	}
}
