package main

import "fmt"

func main() {
	// 单if语句示例
	// var a int = 15
	var a int = 8
	if a < 10 {
		// 当a小于10的时候进入if语句内执行里面的逻辑
		fmt.Println("a小于10")
	}
	// 这里的条件按照顺序结构，后面都是会执行的。
	fmt.Println("a大于10")
}
