package main

import "fmt"

func main() {
	userinfo1 := map[string]string{
		"username": "李四",
		"age":      "21",
		"sex":      "男",
	}
	fmt.Println(userinfo1)
	fmt.Println()

	userinfo2 := map[string]string{
		"username": "张三",
		"age":      "20",
	}

	userinfo2["username"] = "王五"
	fmt.Println(userinfo2)
	fmt.Println()

	userinfo3 := map[string]string{
		"username": "张三",
		"age":      "20",
	}

	fmt.Println(userinfo3["username"])

	v, k := userinfo3["age"]
	fmt.Println(v, k) // 20 true

	z, a := userinfo3["xxx"]
	fmt.Println(z, a) // 空 和 false
	fmt.Println()
	userinfo5 := map[string]string{
		"username": "李四",
		"age":      "21",
		"sex":      "男",
		"height":   "180cm",
	}

	for k, v := range userinfo5 {
		//fmt.Printf("key:%v value:%v\n", k, v)
		fmt.Println(k, v)
	}

	fmt.Println("delete map sex")

	delete(userinfo5, "sex")

	for k, v := range userinfo5 {
		//fmt.Printf("key:%v value:%v\n", k, v)
		fmt.Println(k, v)
	}

}
