package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

	timeNow := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("Time Now:", timeNow)

	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05",
		time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		time.Local))

}
