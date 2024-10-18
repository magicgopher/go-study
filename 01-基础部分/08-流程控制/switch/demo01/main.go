package main

func main() {
	// switch语句示例
	// go语句switch语句改进语法设计，case与case之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行。

	var str = "hello"

	switch str {
	case "hello":
		println(1)
	case "world":
		println(2)
	default:
		println(0)
	}
}
