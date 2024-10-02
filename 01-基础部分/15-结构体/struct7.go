package main

import "fmt"

func main() {
	person := Person{Name: "Alice", Age: 30}
	// UpdatePerson(person)
	// fmt.Println(person.Age) // 输出: 30

	UpdatePerson(&person)   // 指针传递
	fmt.Println(person.Age) // 输出: 31
}

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 定义函数（值传递）
// func UpdatePerson(p Person) {
// 	p.Age += 1 // 只会修改副本
// }

func UpdatePerson(p *Person) {
	p.Age += 1 // 修改原始结构体
}
