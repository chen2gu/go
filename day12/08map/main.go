package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 11
	map1[9] = 134
	map1[2] = 109
	map1[11] = 1100
	map1[41] = 11
	map1[49] = 134
	map1[27] = 109

	// for k, v := range map1 {
	// 	fmt.Println(k, v)
	// }

	var keySlice []int
	for k, _ := range map1 {
		keySlice = append(keySlice, k)

	}
	fmt.Println(keySlice)

	sort.Ints(keySlice)
	fmt.Println(keySlice)

	for _, v := range keySlice {
		fmt.Println(v, map1[v])
	}
}
