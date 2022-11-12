package main

import (
	"machine"
	"time"
)

func main() {
	go blinkExternal()

	blink()
}

func blinkExternal() {
	led := machine.P0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	for {
		led.Low()
		time.Sleep(200 * time.Millisecond)

		led.High()
		time.Sleep(200 * time.Millisecond)
	}
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
