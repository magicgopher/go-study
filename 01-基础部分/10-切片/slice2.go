package main

import "fmt"

func main() {
	// 使用make函数定义一个int类型切片，切片的长度为0，容量为5
	// 此时这个切片是空的，但是这个切片不是nil，空不等同于nil，不等同于nil
	// 空是说明容器是没有东西的，nil是这个容器不存在的
	s1 := make([]int, 0, 5)
	fmt.Println(s1) // 此时的切片是[]

	// 使用append函数向切片中添加元素
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1) // [1 2 3 4]
}
