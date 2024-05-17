package main

import "fmt"

func main() {
	// break带标签示例
	// 第一次循环 外层循环1次(i=0) 内层循环3次(j=0/j=1/j=2)
	// 第二次循环 外层循环1次(i=1) 内层循环3次(j=0/j=1) 此时满足条件 i==1 && j==1 就break outer 跳转到 outer 标签这里执行。
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // break outer 语句会跳出标签为 outer 的外层 for 循环，直接结束整个循环，不再执行内层循环剩余的部分。
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
		fmt.Printf("i: %d\n", i)
	}
	fmt.Println("main end")
}
