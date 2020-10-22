package main

import "fmt"

type calc func(int, int) int

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func testRun() {
	fmt.Println("test Run...")
}

func main() {
	var c calc
	c = add
	fmt.Printf("%T\n", c)
	fmt.Println("ok ?")
	testRun()

	a := 10
	var b myInt = 20
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(a + int(b))
}
