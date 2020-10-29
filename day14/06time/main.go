package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now:", now)

	nowMinute := now.Add(time.Minute)
	fmt.Println("Add Minute:", nowMinute)

	nowHour := now.Add(time.Hour)
	fmt.Println("Add Hour:", nowHour)
}
