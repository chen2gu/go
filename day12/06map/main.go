package main

import "fmt"

func main() {
	userinfo := make(map[string][]string)
	userinfo["username"] = []string{
		"张三",
		"李四",
	} //张三
	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"哈哈",
	}
	fmt.Println(userinfo)

}
