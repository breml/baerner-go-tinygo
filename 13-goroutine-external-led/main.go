package main

import (
	"machine"
	"time"
)

func main() {
	go blinkExternal(machine.P0, 200*time.Millisecond)
	go blinkExternal(machine.P1, 300*time.Millisecond)
	go blinkExternal(machine.P2, 700*time.Millisecond)

	blink(500 * time.Millisecond)
}

func blinkExternal(pin machine.Pin, frequency time.Duration) {
	led := pin
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	for {
		led.Low()
		time.Sleep(frequency)

		led.High()
		time.Sleep(frequency)
	}
}

func blink(frequency time.Duration) {
	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol.Low()

	ledrow := machine.LED_ROW_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		ledrow.Low()
		time.Sleep(frequency)

		ledrow.High()
		time.Sleep(frequency)
	}
}
