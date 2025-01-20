package main

import "fmt"

func main() {
	// 定义一个int数组长度为5，初始化数据只有4个
	var arr5 [5]int = [5]int{1, 2, 3, 4}
	fmt.Printf("arr5的长度:%d\n", len(arr5))
	// 由于arr5数组长度为5，所以arr5[4]没有初始化，所以arr5[4]的值为0【int类型的默认值】
	fmt.Printf("arr5[4]:%d\n", arr5[4])
}
