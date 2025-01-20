package main

import "fmt"

func main() {
	// 不带标签的 continue 示例
	for i := 1; i <= 5; i++ {
		switch i {
		case 1, 3, 5:
			fmt.Printf("Skipping %d\n", i)
			continue // 结束当前循环迭代，直接进入下一次循环
		default:
			fmt.Printf("Processing %d\n", i)
		}
	}
}
