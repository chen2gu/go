package main

import "fmt"

func main() {

	fmt.Println("one")
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"

	fmt.Println(userinfo)
	fmt.Println(userinfo["username"])

	fmt.Println("two")
	var userinfo2 = map[string]string{
		"username": "李四",
		"age":      "21",
		"sex":      "男",
	}

	fmt.Println(userinfo2)

	fmt.Println(userinfo2["username"])

	fmt.Println("therr")
	userinfo3 := map[string]string{
		"username": "王五",
		"age":      "22",
		"sex":      "男",
	}

	fmt.Println(userinfo3)

	fmt.Println(userinfo3["username"])
}
