package main

import "fmt"

var a = "a初始值: 全局变量"

func run() {
	var a = "a初始值: 局部变量"
	var b = "b初始值: 局部变量"
	fmt.Println(a)
	fmt.Println(b)
	return
}

func main() {

	fmt.Println(a)
	run()
	fmt.Println(a)

	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	if flag := true; flag {
		fmt.Println("true")
	}

}
