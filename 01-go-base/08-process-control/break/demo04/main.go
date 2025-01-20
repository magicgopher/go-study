package main

import "fmt"

func main() {
	// 在这个例子中，我们使用了一个带有标签 outer 的 for 循环作为外层循环。
	// 在 switch 语句中，当 i 的值为 2 时，我们使用 break outer 跳出了整个外层的 for 循环，而不仅仅是跳出 switch 语句
outer:
	for i := 0; i < 10; i++ {
		fmt.Println("外层循环迭代:", i)
		switch i {
		case 1:
			fmt.Println("case 1")
		case 2:
			fmt.Println("case 2")
			break outer
		case 3:
			fmt.Println("case 3")
		}
		fmt.Println("外层循环继续")
	}
	fmt.Println("程序结束")
}
