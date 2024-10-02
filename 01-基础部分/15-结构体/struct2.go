package main

import "fmt"

func main() {
	// 初始化结构体
	u1 := User{
		name:    "MagicGopher",
		age:     20,
		sex:     "男",
		address: "地球",
	}
	// 通过 . 的方式获取结构体的字段的值
	fmt.Printf("name:%v\n", u1.name)
	fmt.Printf("age:%v\n", u1.age)
	fmt.Printf("sex:%v\n", u1.sex)
	fmt.Printf("address:%v\n", u1.address)
}

// 定义一个结构体
type User struct {
	name    string
	age     int
	sex     string
	address string
}
