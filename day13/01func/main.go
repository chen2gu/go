package main

import "fmt"

func sunFn(x int, y int) int {
	sum := x + y
	return sum
}

func subFn(x int, y int) int {
	sum := x - y
	return sum
}

func main() {
	sum1 := sunFn(11, 9)
	fmt.Println(sum1)

	sum2 := sunFn(111, 9)
	fmt.Println(sum2)

	sum3 := subFn(111, 9)
	fmt.Println(sum3)
}
