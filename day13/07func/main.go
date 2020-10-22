package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, cb func(int, int) int) int {
	return cb(x, y)
}

func main() {
	sum := calc(2, 3, add)
	Sub := calc(2, 3, sub)
	j := calc(2, 3, func(x, y int) int {
		return x * y
	})
	fmt.Println(sum)
	fmt.Println(Sub)
	fmt.Println(j)

}
