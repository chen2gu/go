package main

import "fmt"

func fn1(x int) {
	x = 11
}

func fn2(x *int) {
	*x = 10
}

func main() {
	a := 1
	fn1(a)
	fmt.Println(a)
	fn2(&a)
	fmt.Println(a)

}
