package main

import "fmt"

func main() {
	// iota定义在一行的情况
	const (
		a, b = iota + 1, iota // iota = 0
		c, d = iota + 1, "哈哈" // iota = 1
		e, f = 55, true       // iota = 2
		g, h = iota, 6.666    // iota = 3
	)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
}
