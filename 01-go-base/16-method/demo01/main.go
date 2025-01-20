package main

import (
	"fmt"
)

func main() {
	/*
		方法：method
			一个方法就是一个包含了接收者的函数，接收者可以是命名类型或者结构体类型的一个值或者是一个指针。
			所有给定类型的方法属于该类型的方法集

		语法：
			func (接收者) 方法名称(参数列表)(返回值列表) {

			}

		总结：method，同函数类似，区别需要有接收者（也就是调用者）

		对比函数：
			A：意义
				方法：某个结构体的行为功能，需要指定某个接收者才能调用。
				函数：一段独立功能的代码，可以直接调用，不需要指定接收者。
			B：语法
				方法：方法名称可以相同。只要接收者不同即可。
				函数：函数名不可以相同。
	*/

	w1 := Worker{
		name: "张三",
		age:  17,
		sex:  "男",
	}
	w1.work()

	w2 := &Worker{
		name: "李四",
		age:  18,
		sex:  "女",
	}
	fmt.Printf("%T\n", w2)
	w2.rest()
}

// 工人结构体
type Worker struct {
	// 字段
	name string
	age  uint
	sex  string
}

// 猫
type Cat struct {
	color string
	age   uint
}

// 定义方法，这里限定接收者（调用者）是Worker类型的结构体
func (w Worker) work() {
	fmt.Println(w.name, "在工作...")
}

// 定义方法，这里限定接收者（调用者）是Worker类型的结构体指针
func (p *Worker) rest() { // p = w2, p = w1的地址
	fmt.Println(p.name, "在休息...")
}

func (p *Worker) printInfo() {
	fmt.Printf("工人姓名:%s, 工人年龄:%d, 工人性别: %s\n", p.name, p.age, p.sex)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫的颜色: %s, 年龄:%d\n", p.color, p.age)
}
