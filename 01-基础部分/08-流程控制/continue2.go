package main

import "fmt"

func main() {
	// 带标签的 continue 示例
outer:
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop iteration: %d\n", i)

		for j := 1; j <= 3; j++ {
			switch j {
			case 2:
				fmt.Printf("Skipping inner loop iteration %d\n", j)
				continue outer // 结束标签为 outer 的循环的当前迭代
			default:
				fmt.Printf("Inner loop processing %d, %d\n", i, j)
			}
		}
	}
}
