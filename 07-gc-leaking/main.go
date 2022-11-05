package main

import (
	"time"
)

func main() {
	i := 0
	for {
		printMemStats()

		x := do(i)
		i += x

		time.Sleep(time.Millisecond * 500)
	}
}

func do(i int) int {
	var buffer [1024]byte
	buffer[i] = 'A' + byte(i)
	return int(buffer[i])
}
