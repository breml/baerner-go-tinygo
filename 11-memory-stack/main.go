// Management of the stack of the goroutines is different, the stack
// can not grow on TinyGo.

package main

const max = 183 // with 184 or higher, there are out of memory panics on microbit V1

var funcs [2]func(int) int

func main() {
	funcs = [2]func(int) int{
		add1, add2,
	}

	res := add(0, 1)
	println(res)
}

func add(sum, a int) int {
	if sum > max {
		return max
	}
	sum += a
	return funcs[a%2](sum)
}

func add1(sum int) int {
	return add(sum, 1)
}

func add2(sum int) int {
	return add(sum, 2)
}
