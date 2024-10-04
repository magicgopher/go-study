package main

import "fmt"

func main() {
	// 定义一个接口类型变量
	var animal Animal

	animal = &Dog{}
	fmt.Println(animal.Speak()) // 输出：汪汪叫

	animal = &Cat{}
	fmt.Println(animal.Speak()) // 输出：喵喵叫
}

// Animal 动物接口
type Animal interface {
	Speak() string // 叫
}

// Dog 狗狗
type Dog struct{}

// Cat 猫猫
type Cat struct{}

// Dog结构体实现Animal接口Speak方法
func (d *Dog) Speak() string {
	return "汪汪叫"
}

// Cat结构体实现Animal接口Speak方法
func (c *Cat) Speak() string {
	return "喵喵叫"
}
