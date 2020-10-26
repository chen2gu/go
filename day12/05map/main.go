package main

import "fmt"

func main() {
	userinfo := make([]map[string]string, 3, 3)
	//fmt.Println(userinfo[0] == nil)

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)

		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["hight"] = "170cm"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)

		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "21"
		userinfo[1]["hight"] = "180cm"
	}

	if userinfo[2] == nil {
		userinfo[2] = make(map[string]string)

		userinfo[2]["username"] = ""
		userinfo[2]["age"] = ""
		userinfo[2]["hight"] = ""
	}

	fmt.Println(userinfo)

	for k, v := range userinfo {
		fmt.Println(k, v)
	}

}
