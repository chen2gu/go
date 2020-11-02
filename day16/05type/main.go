package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Sex   string
	Heiht int
}

func (p Person) PrintInfo() {
	fmt.Printf("Name: %v, Age: %v, Sex: %v\n", p.Name, p.Age, p.Sex)

}

func main() {
	p1 := Person{
		Name:  "Six",
		Age:   0,
		Sex:   "boy",
		Heiht: 0,
	}

	p1.PrintInfo()

	p2 := Person{
		Name:  "Exi",
		Age:   0,
		Sex:   "girl",
		Heiht: 0,
	}

	p2.PrintInfo()

	a := new(int)
	fmt.Printf("%v, %T, %v", a, a, *a)
}
