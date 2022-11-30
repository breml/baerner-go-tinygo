package main

import "time"

const recursions = 33 // for microbit V1, execution time ~10 seconds

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
