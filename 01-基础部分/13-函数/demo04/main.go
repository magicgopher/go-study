package main

import "fmt"

func main() {
	/*
		可变参数的作用
			灵活性：函数可以接受任意数量的同类型参数，调用更加方便。
			代码简洁：不需要预先定义函数的参数个数。
			与切片配合使用：可以方便地将切片传递给可变参数函数。
	*/
	// 调用函数
	add(1, 2, 3, 4, 5)
}

/*
	add()函数受不定数量的参数，参数的类型全部是int类型
	add(1,2,3)
	add(4,5,6,7,8)
*/
// 这种参数写法是可变参数
func add(args ...int) {
	sum := 0
	for _, v := range args {
		sum += v
	}
	fmt.Println("传入的数值之和:", sum)
}
