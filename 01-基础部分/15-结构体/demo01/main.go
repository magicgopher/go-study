package main

import "fmt"

func main() {
	/*
		结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
			通过 struct 关键字来定义结构体
	*/
	// 方式一：使用字段名的方式初始化
	var u1 User
	u1.Name = "MagicGopher1"
	u1.Age = 20
	u1.Sex = "男"
	u1.address = "银河系地球"
	fmt.Println(u1)

	// 方式二：使用键值对方式初始化
	u2 := User{
		Name:    "MagicGopher2",
		Age:     21,
		Sex:     "男",
		address: "银河系地球2",
	}
	fmt.Println(u2)

	// 方式三：使用 new 函数初始化
	p3 := new(User)
	p3.Name = "MagicGopher3"
	p3.Age = 22
	p3.Sex = "男"
	p3.address = "银河系地球3"
	fmt.Println(p3)

	// 方式四：匿名结构体初始化
	p4 := struct {
		Name    string
		Age     int
		Sex     string
		address string
	}{
		Name:    "MagicGopher4",
		Age:     23,
		Sex:     "男",
		address: "银河系地球4",
	}
	fmt.Println(p4)
}

// 定义一个用户结构体
type User struct {
	Name    string
	Age     int
	Sex     string
	address string
}
