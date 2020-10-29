package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixtime := now.Unix()

	fmt.Println("now nuix time :", unixtime)

	unixNatime := now.UnixNano()

	fmt.Println("now Unix Nano time :", unixNatime)
}
