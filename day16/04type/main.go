package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	p1 := Person{
		Name: "What FUACK?",
		Age:  0,
		Sex:  "",
	}

	p2 := p1

	p2.Name = "Test"
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

}
