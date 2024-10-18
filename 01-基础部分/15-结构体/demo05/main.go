package main

import "fmt"

func main() {
	// s1结构体
	s1 := new(Student1)
	s1.name = "学生1"
	s1.year = "2001"
	s1.addr = "地球1"
	s1.score = 101
	// s2结构体
	s2 := new(Student2)
	s2.base.name = "学生2"
	s2.base.year = "2002"
	s2.base.addr = "地球2"
	s2.score = 102
	// 打印s1、s2
	fmt.Println(s1)
	fmt.Println(s2)
}

type Student2 struct {
	base  Base
	score int
}

type Student1 struct {
	Base
	score int
}

type Base struct {
	name string
	year string
	addr string
}
