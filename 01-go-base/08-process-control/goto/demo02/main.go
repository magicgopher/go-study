package main

import "fmt"

func main() {
	// goto: 提供了一种无条件跳转的控制流方式
	i := 0
LOOP:
	for i <= 5 {
		fmt.Println(i)
		i++
		goto LOOP
	}
	fmt.Println("end")
}
