package main

import "fmt"

func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}

	// pass

}

func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}

}

func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}

}

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("---------")
	fn1(5)

	fmt.Println("---------")
	fmt.Println(fn2(100))

	fmt.Println("---------")
	fmt.Println(fn3(3))
}
