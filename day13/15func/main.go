package main

import "fmt"

func fn1() {
	fmt.Println("fn1")
}

func fn2() {
	//fmt.Println("fn2")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()

	panic("error")
}

func fn3() {
	fmt.Println("fn3")
}

func main() {
	fmt.Println("--------->")
	fn1()
	fn2()
	fn3()
}
