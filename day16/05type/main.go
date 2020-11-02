package main

import "fmt"

func main() {
	a := new(int)
	fmt.Printf("%v, %T, %v", a, a, *a)
}
