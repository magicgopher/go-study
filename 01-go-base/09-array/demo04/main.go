package main

import "fmt"

func main() {
	// 定义数组时，数组的长度可以使用...替代，这样数组的长度就可以根据元素的个数自动计算。
	var arr4 = [...]int{10, 20, 30, 40, 50, 60}
	fmt.Printf("arr4的长度:%d\n", len(arr4))

	// 访问数组不存在的索引
	// fmt.Printf("arr4[6]:%d\n", arr4[6]) // invalid argument: index 6 out of bounds [0:6]
}
