package main

import "fmt"

type Persion struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Persion
	p1.name = "sam"
	p1.age = 11
	p1.sex = "boy"

	fmt.Printf("%v,%T\n", p1, p1)
	fmt.Println("------->")

	fmt.Printf("%#v,%T", p1, p1)

}
