package main

import "fmt"

func main() {
	// 嵌套if else示例
	age := 25
	height := 180

	if age >= 18 {
		fmt.Println("年龄符合要求")
		if height >= 160 {
			fmt.Println("身高符合要求")
			fmt.Println("符合报名条件")
		} else {
			fmt.Println("身高不符合要求")
			fmt.Println("不符合报名条件")
		}
	} else {
		fmt.Println("年龄不符合要求")
		fmt.Println("不符合报名条件")
	}
}
