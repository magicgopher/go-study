package main

import "fmt"

// 全局变量
var counter int // 默认值：0

func main() {
	counter++
	counter++
	fmt.Println(counter) // 输出结果：2
}
