package main

import (
	"fmt"
	"time"
)

func zdss() {
	for {
		time.Sleep(time.Second)
		fmt.Println("Run.....")
	}
}

func main() {
	//fmt.Println("q")
	n := 5
	for t := range time.NewTicker(time.Second).C {
		time.Sleep(time.Second * 2)
		n--
		fmt.Println(t)
		//time.Sleep(time.Second)
		if n == 0 {
			time.NewTicker(time.Second).Stop()
			fmt.Println("end.")
			break
		}

	}
}

// test
