package main

import (
	"machine"
	"time"
)

func main() {
	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol.Low()

	ledrow := machine.LED_ROW_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		ledrow.Low()
		println("low")
		time.Sleep(500 * time.Millisecond)

		ledrow.High()
		println("high")
		time.Sleep(500 * time.Millisecond)
	}
}
