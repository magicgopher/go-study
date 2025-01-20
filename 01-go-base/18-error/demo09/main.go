package main

import (
	"fmt"
	"log"
)

// 自定义错误类型
type BusinessError struct {
	Code    int
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("业务错误 [%d]: %s", e.Code, e.Message)
}

// 包装recover的工具函数
func SafelyDo(work func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("检测到panic: %v\n", err)
			// 这里可以添加报警或监控代码
		}
	}()

	work()
}

func main() {
	// 示例1：处理自定义业务错误
	SafelyDo(func() {
		processBusinessLogic("")
	})

	// 示例2：处理运行时错误
	SafelyDo(func() {
		var slice []int
		slice[0] = 1 // 将触发panic
	})

	fmt.Println("程序继续执行...")
}
