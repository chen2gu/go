package main

import "fmt"

func cacl(index string, a, b int) int {
	ret := a + b

	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2

	defer cacl("AA", x, cacl("A", x, y))
	x = 10

	defer cacl("AA", x, cacl("A", x, y))
	y = 20
}
