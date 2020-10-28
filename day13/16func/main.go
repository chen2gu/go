package main

import "fmt"

func fn1() {
	fmt.Println("1.--------->")
}

func fn2(a, b int) int {
	//fmt.Println("2.--------->")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()

	return a / b
}
func fn3() {
	fmt.Println("3.--------->")
}

func main() {
	fmt.Println("0.--------->")
	fn1()
	fmt.Println(10, 0)
	fmt.Println(fn2(10, 0))
	fn3()
	fmt.Println(fn2(10, 2))

}
