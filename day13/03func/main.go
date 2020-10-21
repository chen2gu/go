package main

import "fmt"

// int 升序排序
func sortInAsc(slice []int) []int {
	//
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice

}

func main() {

	slinA := []int{12, 33, 1, 56, 85, 22, 19, 11, 1, 0, -1}
	arr := sortInAsc(slinA)
	fmt.Println(arr)
	fmt.Println("Ok ?")
}
