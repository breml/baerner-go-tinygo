// programm using global memory in .bss and .data
// printing memory statistics
//
// quite some trickery is necessary to prevent that the compiler does not
// optimize the memory locations away.

package main

import (
	"runtime"
	"time"
)

const size = 1024

var bbsBlob [size]byte

var dataBlob [size]byte = [size]byte{1, 2}

func main() {
	for i := 0; i < size; i++ {
		bbsBlob[i] = byte(time.Now().Nanosecond() % 26)
		dataBlob[i] = byte(time.Now().Nanosecond() % 26)
	}
	item := time.Now().Nanosecond() % (size - 1)

	println(bbsBlob[item]+'A', " (prevent compiler optimization)")
	println(dataBlob[item]+'A', " (prevent compiler optimization)")

	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	println("GCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)
}
