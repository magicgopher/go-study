package main

import "fmt"

func main() {
	// 结构体是引用类型（地址拷贝）、而数组是值类型（值拷贝）
	// 初始化结构体
	u1 := User{name: "MagicGopher", age: 20}
	fmt.Println(u1)
	fmt.Printf("%p, %T\n", &u1, u1)

	u2 := u1
	fmt.Println(u2)
	fmt.Printf("%p, %T\n", &u2, u2)

	// 修改u2的数据
	u2.name = "Martin Garrix"
	u2.age = 28

	// 再输出u1、u2
	fmt.Println(u1)
	fmt.Println(u2)

	u3 := &u1
	fmt.Println(u2)
	fmt.Printf("%p, %T\n", &u3, u3)

	// 修改 u1
	u1.name = "haha"

	// 在输出u1、u3
	fmt.Println(u1)
	fmt.Println(*u3)
}

type User struct {
	name string
	age  int
}
