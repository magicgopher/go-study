package main

import "fmt"

// 单个返回值
func add(a, b int) int {
	return a + b
}

// 多个返回值
func division(a, b float64) (float64, float64) {
	return a + b, a - b
}

// 命名返回值
func swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

// 使用空白标识符
func getCircleProps(radius float64) (area float64, perimeter float64) {
	area = 3.14 * radius * radius
	perimeter = 2 * 3.14 * radius
	return
}

func main() {
	// 单个返回值
	result := add(5, 3)
	fmt.Println("函数单个返回值:", result)

	// 多个返回值
	quotient, remainder := division(10.0, 3.0)
	fmt.Println("函数多个返回值:", quotient, remainder)

	// 命名返回值
	x, y := swap(10, 20)
	fmt.Println("命名返回值:", x, y)

	// 使用空白标识符（忽略返回值）
	area, _ := getCircleProps(5.0)
	fmt.Println("Circle properties:", area)
}
