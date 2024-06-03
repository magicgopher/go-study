package main

import "fmt"

func main() {
	// 创建map
	m1 := map[int]string{
		1: "java",
		2: "golang",
		3: "python",
		4: "c++",
		5: "rust",
	}
	// 通过键获取值
	fmt.Printf("m1[1] = %v\n", m1[3])
	// 获取不存在的键，返回值是数据类型的默认值
	fmt.Printf("m1[6] = %v\n", m1[6])
}
