package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程
			支持将一个函数作为另一函数的参数（回调函数）
			也支持将一个函数作为另一个函数的返回值（闭包）

		闭包（closure）
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），并且该外层函数的返回值就是这个内层函数。
			这个内层函数和外层函数的局部变量，统称为闭包结构。
	*/

	/*
		结果分析：
			res1 是一个函数，每次调用它，i 都会增加并返回当前的值。
			v1 := res1()，第一次调用 res1，i 变为 1，输出 1。
			v2 := res1()，第二次调用 res1，i 变为 2，输出 2。
			res2 := increment()，再次调用 increment() 函数，此时返回的是一个新的函数，而不是之前的 res1。
			v3 := res2()，调用新返回的函数，此时它从新的起始点开始计数，所以输出 1。

			因此，v3 不等于 3 是因为 res2 是一个新的闭包，它的内部变量 i 是独立于之前的 increment() 返回的函数的。
			每次调用 increment() 都会创建一个新的闭包环境，因此它们的状态是相互独立的。
	*/

	// 调用
	res1 := increment() // 这里的res1就是increment函数内的fun
	fmt.Printf("%T\n", res1)
	fmt.Println(res1) // 这里输出的函数的内存地址

	v1 := res1()
	fmt.Println(v1)
	v2 := res1()
	fmt.Println(v2)

	// 这里又再次调用了increment函数
	res2 := increment()
	fmt.Println(res2)
	v3 := res2()
	fmt.Println(v3)
}

func increment() func() int { // 外层函数
	// 定义一个局部变量
	i := 0
	// 定义一个匿名函数，将函数赋值给变量，将其返回
	fun := func() int { // 内层函数
		i++
		return i
	}
	// 返回该匿名函数
	return fun
}
