package main

import "fmt"

func adder() func() int {
	i := 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	i := 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {

	fn := adder()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fmt.Println("--------->")
	fn2 := adder2()
	fmt.Println(fn2(1))
	fmt.Println(fn2(2))
	fmt.Println(fn2(3))
}
