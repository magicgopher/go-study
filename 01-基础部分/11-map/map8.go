package main

import "fmt"

func main() {
	// 定义一个map
	myMap := map[int]string{
		1: "apple",
		2: "banana",
		3: "orange",
	}
	// 判断键的值是否存在
	value, ok := myMap[1]
	fmt.Printf("value:%v, ok:%v\n", value, ok)
	value, ok = myMap[4]
	fmt.Printf("value:%v, ok:%v\n", value, ok)
}
