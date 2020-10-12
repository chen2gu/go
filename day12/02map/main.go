package main

import "fmt"

func main() {
	userinfo := map[string]string{
		"username": "李四",
		"age":      "21",
		"sex":      "男",
	}

	for k, v := range userinfo {
		fmt.Printf("key:%v value:%v\n", k, v)
		//fmt.Println(k, v)
	}

}
