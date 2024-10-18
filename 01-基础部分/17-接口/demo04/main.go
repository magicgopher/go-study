package main

import "fmt"

func main() {
	var a any
	a = 12
	if v, ok := a.(int); ok {
		fmt.Println(v)
	}
}
