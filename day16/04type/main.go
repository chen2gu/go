package main

import "fmt"

func main() {
	a := make([]int, 8, 8)
	fmt.Printf("%v, %T", a, a)
}
