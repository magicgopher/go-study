package main

import "fmt"

func main() {
	/*
		Go语言不是纯面向对象的语言
		继承：可以通过结构体嵌套的方式模拟（匿名字段）
		多态：可以通过接口的方式来模拟

		多态：一个事物的多种形态
		go语言通过接口模拟多态

		就是一个接口的实现
			1.看成实现本身的实现，能够访问实现类中的属性和方法
			2.看成是对应的接口类型，那就只能够访问接口中的方法
	*/

	// 创建不同类型的动物实例
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "小花"}

	// 调用函数，传入不同类型的动物实例
	LetAnimalSpeak(dog)
	LetAnimalSpeak(cat)
}

// 定义一个接口
type Animal interface {
	Speak() string
}

// 定义不同类型的动物结构体
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// 实现接口方法
func (d Dog) Speak() string {
	return "汪汪汪！"
}

func (c Cat) Speak() string {
	return "喵喵喵！"
}

// 接受接口类型作为参数的函数
func LetAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}
