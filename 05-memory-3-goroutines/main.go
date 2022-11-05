package main

import (
	"sync"
	"time"
)

// max value
//   microbit: 5
//   microbit v2: 58
const max = 5

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(max)
	for i := 0; i < max; i++ {
		go wait(time.Duration(100+i)*time.Millisecond, wg)
	}
	println("wait")
	wg.Wait()
	println("done")
}

func wait(d time.Duration, wg *sync.WaitGroup) {
	time.Sleep(d)
	wg.Done()
}
