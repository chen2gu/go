package main

import "fmt"

func main() {
	a := 10
	p := &a

	//fmt.Printf(("a: ",a)
	fmt.Printf("a: %v,%T,%p\n", a, a, &a)
	fmt.Printf("p: %v,%T,%p", p, p, &p)

}
