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
	(*v2).name = "fire"
	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v2, v2)

	v3 := Person{
		name: "six",
		age:  022,
		city: "SZ",
	}

	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v3, v3)

	v4 := &Person{
		name: "zhaosi",
		age:  12,
		city: "CN",
	}
	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v4, v4)

	v5 := &Person{
		city: "SZ",
	}
	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v5, v5)

	v6 := &Person{
		"zhaosi",
		11,
		"SZ",
	}
	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v6, v6)

	var v7 = &Person{
		name: "zhaosi",
		age:  112,
		city: "CN",
	}

	fmt.Println("------->")
	fmt.Printf("%#v,     %T\n", v7, v7)
}
