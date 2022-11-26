// "empty" program, just reporting memory statistics on serial as a base line.

package main

import "runtime"

func main() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	// GCSys: memory used by the garbage collector, 32 byte base + 2 bit per 16 byte block
	// HeapIdle: available heap memory
	// HeapInuse: used heap memory
	// HeapReleased: heap memory released to the OS again (does not happen in TinyGo)
	// HeapSys: HeapIdle + HeapInUse
	// Sys: total memory available (~ heapSys + GCSys rounded to 16 byte blocks)
	println("GCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)
}
