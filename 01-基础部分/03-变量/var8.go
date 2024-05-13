package main

func main() {
	// Go语言变量注意事项7
	// 变量定义了就要使用，否则无法通过编译。
	// var x int
	// 定义了变量 x，但未使用它

	var y string = "Hello, world!"
	// 定义了变量 y，并使用了它

	_ = y // 使用了变量 y，但是将其赋值给了匿名变量 _

	// 此处未使用变量 x，将会导致编译错误
}
