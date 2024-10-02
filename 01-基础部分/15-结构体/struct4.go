package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// fun1()
	fun2()
}

func fun1() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := p1 // 这里是值拷贝

	p2.Name = "Bob" // 修改 p2 不会影响 p1

	fmt.Println(p1.Name) // 输出: Alice
	fmt.Println(p2.Name) // 输出: Bob
}

func fun2() {
	p1 := &Person{Name: "Alice", Age: 30} // p1 是指向结构体的指针
	p2 := p1                              // 这里是指针拷贝

	p2.Name = "Bob" // 修改 p2 会影响 p1

	fmt.Println(p1.Name) // 输出: Bob
	fmt.Println(p2.Name) // 输出: Bob
}
