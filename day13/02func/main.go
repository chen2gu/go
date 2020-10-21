package main

import "fmt"

func sunFn(x, y int) int {
	return x + y
}

func cacl(z, c int) (int, int) {
	sum := z + c
	sub := z - c
	return sum, sub
}

func cacl1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum1 := sunFn(11, 12)
	fmt.Println(sum1)

	fmt.Println("---------")

	a, b := cacl(10, 2)
	fmt.Println(a, b)

	fmt.Println("---------")
	a, b = cacl1(15, 2)
	fmt.Println(a, b)

}
