package main

// Maximum without error
//   microbit: 29
//   microbit-v2: 26
//   notebook: 510
//   notebook (regular Go): 9586957
const max = 26

var c chan func() int

func main() {
	c = make(chan func() int, 0)

	go func() {
		for i := 0; i < max; i++ {
			c <- sum
		}
		close(c)
	}()

	res := sum()
	println(res)
}

func sum() int {
	res := 1
	for f := range c {
		res += f()
	}
	return res
}
