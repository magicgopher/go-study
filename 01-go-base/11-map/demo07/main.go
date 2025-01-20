package main

import "fmt"

func main() {
	// 使用delete删除键值对
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	// for range遍历map
	for k, v := range m {
		println(k, v)
	}
	// 删除键值对
	delete(m, "a")
	fmt.Println("============")
	// 再遍历map
	for k, v := range m {
		println(k, v)
	}
	// 追加元素
	m["test"] = 100
	fmt.Println("============")
	// 再遍历map
	for k, v := range m {
		println(k, v)
	}
}
