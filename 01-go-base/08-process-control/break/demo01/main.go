package main

import "fmt"

func main() {
	// break不带标签示例
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // 当 i 等于 3 时，跳出当前 for 循环
		}
		fmt.Println(i) // 这里只会打印1到4
	}
}
