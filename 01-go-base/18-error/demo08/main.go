package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建几个错误
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	// 使用 Join 函数将错误合并
	combinedErr := errors.Join(err1, err2, err3)

	// 打印合并后的错误
	fmt.Println(combinedErr)
}
