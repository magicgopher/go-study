package main

import "fmt"

// 定义一个嵌入的结构体
type Address struct {
	City  string
	State string
}

// 定义一个包含匿名字段的结构体
type Person struct {
	Name    string
	Address // 匿名字段
}

func main() {
	p := Person{
		Name: "Alice",
		Address: Address{
			City:  "Wonderland",
			State: "Fantasy",
		},
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.City) // 直接访问匿名字段的字段
	fmt.Println("State:", p.State)
}
