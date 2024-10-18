package main

import (
	"fmt"
)

// 定义一个USB接口
type USB interface {
	Connect()
	Disconnect()
}

// 定义鼠标结构体
type Mouse struct {
	Name string
}

// 鼠标结构体实现USB接口的Connect方法
func (m Mouse) Connect() {
	fmt.Println("Mouse connected.")
}

// 鼠标结构体实现USB接口的Disconnect方法
func (m Mouse) Disconnect() {
	fmt.Println("Mouse disconnected.")
}

func main() {
	// 创建一个Mouse对象
	mouse := Mouse{Name: "Logitech"}

	// 使用USB接口类型的变量
	var usbDevice USB
	usbDevice = mouse // 将Mouse对象赋值给USB类型的变量

	// 调用USB接口方法
	usbDevice.Connect()    // 输出：Mouse connected.
	usbDevice.Disconnect() // 输出：Mouse disconnected.
}
