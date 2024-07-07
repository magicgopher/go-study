package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("初始切片:", numbers)

	growSlice(&numbers, 3)
	fmt.Println("修改长度后切片:", numbers)
}

func growSlice(s *[]int, newLength int) {
	*s = (*s)[0:newLength]
}
