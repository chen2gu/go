package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {
	// sliceKey []string
	var sliceKey []string
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)

	}

	var str string

	sort.Strings(sliceKey)
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>>%v\n", v, map1[v])

	}
	return str
}

func main() {
	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "18",
		"sex":      "男",
		"height":   "180",
		"sasa":     "xxx",
	}

	fmt.Println(m1)
	str := mapSort(m1)
	fmt.Printf(str)
}
