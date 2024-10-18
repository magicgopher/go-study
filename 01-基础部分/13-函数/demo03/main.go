package main

import (
	"fmt"
)

func main() {
	// 调用函数
	// 1-10的和
	getSum(10)
	// 1-50的和
	getSum(50)
	// 1-100的和
	getSum(100)
}

// getSum 定义一个范围求和函数
// 参数n：求和的范围，传入10就是求1-10的和,传入50就是求1-50的和
func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是: %d\n", n, sum)
}
