package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()

	fmt.Println(timeObj)

	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
	var str = timeObj.Format("2006-01-02 15:04:05")
	// wath 29290-10-10 11:488 funck
	fmt.Println(str)
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

	//fmt.Println(now)

}
