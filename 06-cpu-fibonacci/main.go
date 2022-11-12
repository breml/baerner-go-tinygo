package main

import "time"

// For ~ 10 sec runtime:
//   microbit: 33
//   microbit-v2: 35
//   notebook: 47 (go versus tinygo)
const recursions = 33

func main() {
	print("result: ")
	start := time.Now()
	res := FibonacciRecursion(recursions)
	println(res)
	println(time.Since(start).String())
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
