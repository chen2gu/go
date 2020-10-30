package main

import "fmt"

func main() {
	a := 5
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println("------->")

	userinfo := make(map[string]string)
	userinfo["userinfo"] = "shanshan"
	fmt.Println(userinfo)

	fmt.Println("------->")
	var b = make([]int, 5, 5)
	fmt.Println(b)

	fmt.Println("------->")

	c := new(int)
	d := new(bool)
	fmt.Println(*c)
	fmt.Println(*d)
}
