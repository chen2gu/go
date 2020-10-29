package main

import (
	"fmt"
	"time"
)

func main() {
	unixTime := time.Now().Unix()
	fmt.Println("unix Time:", unixTime)

	fmt.Println(time.Unix(time.Now().Unix(), 0))
	//buh
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

}
