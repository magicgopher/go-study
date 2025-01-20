package main

import "fmt"

func main() {
	// for range循环示例。
	// 遍历数组。
	arr := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("%T\n", arr)
	for i, v := range arr {
		println(i, v)
	}

	fmt.Println("=========================")

	// 遍历切片。
	slice := []int{100, 200, 300, 400, 500}
	fmt.Printf("%T\n", slice)
	for i, v := range slice {
		println(i, v)
	}

	fmt.Println("=========================")

	// 遍历map
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Printf("%T\n", m)
	for k, v := range m {
		println(k, v)
	}

	fmt.Println("=========================")

	// 遍历字符串
	str := "hello world"
	fmt.Printf("%T\n", str)
	for i, v := range str {
		println(i, v)
	}
}
