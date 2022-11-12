// "empty" program, just reporting memory statistics on serial as a base line.

package main

import "runtime"

func main() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	println("GCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)
}
