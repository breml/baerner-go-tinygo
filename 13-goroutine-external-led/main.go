package main

import (
	"machine"
	"time"
)

func main() {
	go blinkExternal(machine.P1, 200*time.Millisecond)
	go blinkExternal(machine.P2, 300*time.Millisecond)
	go blinkExternal(machine.P5, 700*time.Millisecond)

	blink()
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

func blink() {
	rotarySensor := adc{Pin: machine.P0}

	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol.Low()

	ledrow := machine.LED_ROW_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		frequency := time.Duration(rotarySensor.Get() >> 6)

		ledrow.Low()
		time.Sleep(frequency * time.Millisecond)

		frequency = time.Duration(rotarySensor.Get() >> 6)

		ledrow.High()
		time.Sleep(frequency * time.Millisecond)
	}
}
