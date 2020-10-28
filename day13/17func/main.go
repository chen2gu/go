package main

import (
	"errors"
	"fmt"
)

func readFile(filename string) error {
	if filename == "main.go" {
		return nil
	} else {
		return errors.New("read error.")
	}

}
func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("myFn:", err)
		}
	}()

	err := readFile("xx.go")
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("---------->")
	myFn()
}
