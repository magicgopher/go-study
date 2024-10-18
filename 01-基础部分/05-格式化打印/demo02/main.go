package main

import "fmt"

// User 用户结构体
type User struct {
	// 姓名
	Name string
	// 年龄
	Age int
}

func main() {
	// Printf 格式化输出
	// func Printf(format string, a ...any) (n int, err error)

	// 格式化打印输出通用的占位符号
	str1 := "Hello Google"
	// %v 占位符
	fmt.Printf("str1:%v\n", str1) // str1:Hello Google

	// %+v 占位符
	user := User{
		Name: "MagicGopher",
		Age:  20,
	}
	fmt.Printf("user:%+v\n", user) // user:{Name:MagicGopher Age:20}
	// %#v 占位符
	fmt.Printf("user:%#v\n", user) // user:main.User{Name:"MagicGopher", Age:20}

	f1 := 6.666
	str2 := "Hello Golang"
	// %T 占位符
	fmt.Printf("f1数据类型:%T\n", f1)     // f1数据类型:float64
	fmt.Printf("str2数据类型:%T\n", str2) // str2数据类型:string

	// 在Printf格式化打印输出要显示%占位符时，需要使用%%占位符
	b1 := true
	fmt.Printf("b1 %%T输出结果:%T\n", b1) // b1 %T输出结果:bool
}
