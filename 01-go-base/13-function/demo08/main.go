package main

import "fmt"

func main() {
	// 打印一下func1函数的数据类型
	fmt.Printf("%T\n", func1) // 结果：func(int, int)
	fmt.Println(func1)        // 这里输出的是函数的内存地址：0x104436db0

	// 定义一个函数类型的变量
	var f1 func(int, int)
	fmt.Println(f1) // <nil> 空指针未分配状态

	// 给 f1 函数变量赋值，这里必须保证赋值的数据和声明的数据类型是一致的
	f1 = func1      // 将func1的值(函数体的地址)赋值给f1
	fmt.Println(f1) // 这里f1输出的就是func1的地址：0x104436db0

	// 调用func1函数
	func1(10, 20)
	// 调用f1函数变量
	f1(11, 22)

	res1 := func2         // 将func2的值(函数的内存地址)赋值给res1，res1和func2指向同一个函数体
	res2 := func2(33, 55) // 将func2函数进行调用，将函数的执行结果赋值给res2，相当于：a+b
	fmt.Println(res1)
	fmt.Println(res2)
	// 打印类型
	fmt.Printf("res1类型:%T\n", res1)
	fmt.Printf("res2类型:%T\n", res2)

	fmt.Println(res1(1, 5)) // 赋值的函数变量，也可以通过函数变量名()的方式来使用
	// res2() // 这里赋值之后是 int类型，所以不能将其当成函数来使用
}

// 定义一个func1函数
func func1(a, b int) {
	fmt.Printf("a:%d, b:%d\n", a, b)
}

// 再定义一个func2函数
func func2(a, b int) int {
	return a + b
}
