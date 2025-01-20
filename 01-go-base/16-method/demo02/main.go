package main

import (
	"fmt"
)

func main() {
	/*
		OOP中的继承性：
		如果两个类（class）存在继承关系，其中一个是子类，另一个作为父类，那么：

		1.子类可以直接访问父类的属性和方法
		2.子类可以新增自己的属性和方法
		3.子类可以重写父类的方法（override就是将父类已有的方法，重新实现）

		Go语言通过结构体嵌套：
			1，模拟继承性：is - a
			type A struct {
				field
			}
			type A struct {
				A // 匿名字段
			}
			2.模拟聚合关系：has - a
			type C struct {
				field
			}
			type D struct {
				c C // 聚合关系
			}
	*/

	// 创建一个Person
	p1 := Person{name: "张三", age: 18}
	fmt.Println(p1.name, p1.age) // 打印父类属性
	p1.eat()                     // 调用方法

	// 创建Student
	s1 := Student{Person{name: "李四", age: 20}, "SCUT"}
	fmt.Println(s1.name, s1.age) // 打印子类的属性（子类访问父类的定义的字段，其实就是提升字段）
	fmt.Println(s1.school)       // 子类独有的字段属性
	s1.eat()                     // 调用eat方法（没有重写调用的就是父类的，重写了就是调用子类的）
	s1.study()                   // 子类对象访问自己的方法
	s1.eat()                     // 再次调用eat方法
	s1.Person.eat()              // 调用父类的eat

	// 创建Worker
	w1 := Worker{Person{name: "王五", age: 21}, 10000}
	fmt.Println(w1.name, w1.age) // 打印子类的属性（子类访问父类的定义的字段，其实就是提升字段）
	fmt.Println(w1.salary)       // 子类独有的字段属性
	w1.eat()                     // 调用eat方法（没有重写调用的就是父类的，重写了就是调用子类的）
	w1.work()                    // 子类对象访问自己的方法
	w1.eat()                     // 再次调用eat方法
	w1.Person.eat()              // 调用父类的eat
}

// 定义一个父类（父结构体）
type Person struct {
	name string
	age  uint
}

// 定义一个子类 Student
type Student struct {
	Person        // 结构体嵌套，模拟继承性（匿名结构体字段）
	school string // 子类独有的字段属性
}

// 定义一个子类 Worker
type Worker struct {
	Person         // 结构体嵌套，模拟继承性（匿名结构体字段）
	salary float64 // 子类独有的字段属性
}

// 父类的方法
func (p Person) eat() {
	fmt.Println("父类的方法，吃饭方法...")
}

// 子类的方法 study 接收者：Student
func (s Student) study() {
	fmt.Println("子类的方法，学生学习...")
}

// 子类重写的方法 接收者 Student 重写了父结构体的eat方法
func (s Student) eat() {
	fmt.Println("子类Student重写的eat方法...")
}

// 子类的方法 work 接收者：Worker
func (w Worker) work() {
	fmt.Println("子类的方法，在工作中...")
}

// 子类重写的方法 接收者 Worker 重写了父结构体的eat方法
func (w Worker) eat() {
	fmt.Println("子类Worker重写的eat方法...")
}
