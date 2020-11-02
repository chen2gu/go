package main

import "fmt"

type Person struct {
	name string
	age  int
	city string
}

func main() {

	v2 := new(Person)
	v2.name = "haha"
	v2.age = 11
	v2.city = "xxxx"

	fmt.Printf("------->%#v,%T", v2, v2)

}
