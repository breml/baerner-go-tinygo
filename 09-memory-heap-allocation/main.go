package main

import (
	"runtime"
	"time"
)

func main() {
	// memory allocated on the heap
	var buffer [11 * 1024]byte

	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	println("\nGCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)

	i := 0
	for {
		buffer[i] = 'A'
		i++

		time.Sleep(100 * time.Millisecond)
		if i == len(buffer) {
			break
		}
	}
	print(string(buffer[0]), string(buffer[len(buffer)-1]))
}
