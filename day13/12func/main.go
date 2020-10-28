package main

import "fmt"

func f1() {
	fmt.Println("start--------->")
	defer func() {
		fmt.Println("stop0--------->")
		fmt.Println("stop1--------->")
	}()

	fmt.Println("end--------->")
}

func f2() int {
	var a int

	defer func() {
		a++
	}()
	return a
}

func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func main() {

	f1()

	fmt.Println("--------->")
	fmt.Println(f2())

	fmt.Println("--------->")
	fmt.Println(f3())

}
