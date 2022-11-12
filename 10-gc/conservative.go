//go:build gc.conservative
// +build gc.conservative

package main

import "runtime"

func printMemStats() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	println("GCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)
}
