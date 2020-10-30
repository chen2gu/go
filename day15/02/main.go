package main

import "fmt"

func main() {
	a := 11
	p := &a

	//
	//fmt.Printf("a: %v,%T,%p\n",a,a,&a)
	//fmt.Printf("p: %v,%T,%p",p,p,&p)
	fmt.Println(p)  //di zhi
	fmt.Println(*p) //zhi

	*p = 30

	fmt.Println(a)
}
