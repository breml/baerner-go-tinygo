// https://tinygo.org/docs/concepts/compiler-internals/interrupts/

package main

import (
	"machine"
	"time"
)

func main() {
	start := time.Now()

	led := machine.P0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	machine.BUTTONA.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
		if p.Get() {
			led.Low()
			println("Button A up after", time.Since(start).Round(time.Millisecond).String())
		} else {
			led.High()
			println("Button A down after", time.Since(start).Round(time.Millisecond).String())
		}
	})

	machine.BUTTONB.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
		println("Button B pressed after", time.Since(start).Round(time.Millisecond).String())
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
