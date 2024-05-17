package main

func main() {
	// 使用switch语句判断DAY日期是星期几
	DAY := 6
	switch DAY {
	case 1:
		println("星期一")
	case 2:
		println("星期二")
	case 3:
		println("星期三")
	case 4:
		println("星期四")
	case 5:
		println("星期五")
	case 6:
		println("星期六")
	case 7:
		println("星期日")
	default:
		println("输入错误")
	}
}
