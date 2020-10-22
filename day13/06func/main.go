package main

import "fmt"

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calc
	c = add
	fmt.Printf("%T\n", c)
	fmt.Println("ok ?")
}
