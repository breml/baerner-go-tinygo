package main

import (
	"runtime"
	"sync"
	"time"
)

const max = 5 // maximum for microbit V1

var m runtime.MemStats

func main() {
	printMemStats()

	wg := &sync.WaitGroup{}

	wg.Add(max)
	for i := 0; i < max; i++ {
		go wait(time.Duration(100+i)*time.Millisecond, wg)
		printMemStats()
	}
	println("wait")
	wg.Wait()
	println("done")

	printMemStats()
	runtime.GC()

	printMemStats()
}

func wait(d time.Duration, wg *sync.WaitGroup) {
	time.Sleep(d)
	wg.Done()
}

func printMemStats() {
	runtime.ReadMemStats(&m)
	println("GCSys: ", m.GCSys, " | HeapIdle: ", m.HeapIdle, " | HeapInuse: ", m.HeapInuse, " | HeapReleased: ", m.HeapReleased, " | HeapSys: ", m.HeapSys, " | Sys: ", m.Sys)
}
