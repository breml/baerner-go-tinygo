package main

import (
	"time"
)

func main() {
	start := time.Now()

	write := func(text string, frequency time.Duration) {
		for {
			println(time.Since(start).Round(time.Millisecond).String(), text)
			time.Sleep(frequency)
		}
	}

	go write("foo", 2000*time.Millisecond)

	write("bar", 900*time.Millisecond)
}
