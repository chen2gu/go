package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "how do you do Cannot use ss(type 		 k8sClient.ResourceN in map index )"

	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	var strMap = make(map[string]int)

	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
