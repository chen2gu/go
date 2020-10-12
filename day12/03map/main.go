package main

import "fmt"

func main() {
	userinfo1 := map[string]string{
		"username": "李四",
		"age":      "21",
		"sex":      "男",
	}
	fmt.Println(userinfo1)

	userinfo2 := map[string]string{
		"username": "张三",
		"age":      "20",
	}

	userinfo2["username"] = "王五"
	fmt.Println(userinfo2)

	userinfo3 := map[string]string{
		"username": "张三",
		"age":      "20",
	}

	fmt.Println(userinfo3["username"])
}
