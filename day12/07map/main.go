package main

import "fmt"

func main() {
	userinfo1 := make(map[string]string)
	userinfo1["username"] = "zhang san"
	userinfo1["age"] = "20"

	userinfo2 := userinfo1
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)

}
