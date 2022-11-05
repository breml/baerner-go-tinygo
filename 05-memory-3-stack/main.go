package main

// Max value for:
//   microbit: 183 (184, 186)
const max = 183

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
