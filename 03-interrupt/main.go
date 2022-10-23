package main

import (
	"machine"
	"time"
)

func main() {
	start := time.Now()

	machine.BUTTONA.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
		println("Button A pressed after", time.Since(start).Round(time.Millisecond).String())
	})

	blink()
}

func blink() {
	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol.Low()

	ledrow := machine.LED_ROW_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		ledrow.Low()
		time.Sleep(500 * time.Millisecond)

		ledrow.High()
		time.Sleep(500 * time.Millisecond)
	}
}
