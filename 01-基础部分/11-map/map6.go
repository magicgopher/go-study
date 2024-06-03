package main

import "fmt"

func main() {
	// 可以用 for range 获取 key-value 对，key 是只读的，value 是可读写的。
	// Go 语言中 map 的 key 是不能直接修改的。
	// 定义map
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("修改前的map:", m)
	// 遍历 map，修改 value
	for k, v := range m {
		fmt.Println("修改前:", k, v)
		m[k] = v * 2
		fmt.Println("修改后:", k, m[k])
	}
	fmt.Println("修改后的map:", m)
}
