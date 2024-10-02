package main

import (
	"fmt"
)

func main() {
	// 创建mouse类型
	m1 := Mouse{"罗技鼠标"}
	fmt.Println(m1.name)
	// 创建FlashDisk类型
	f1 := FlashDisk{"金士顿"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)

	var usb USB
	// usb = m1 // 这里的usb接口的类型就是m1 Mouse结构体的实例
	usb = f1 // 这里的usb接口的类型就是f1 FlashDisk结构体的实例
	usb.start()
	usb.end()

	f1.deleteData()
}

// 定义接口
type USB interface {
	start() // 开始工作
	end()   // 结束工作
}

// 鼠标结构体
type Mouse struct {
	name string
}

// u盘
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("鼠标开始工作...")
}

func (m Mouse) end() {
	fmt.Println("鼠标结束工作...")
}

func (f FlashDisk) start() {
	fmt.Println("U盘开始工作...")
}

func (f FlashDisk) end() {
	fmt.Println("U盘结束工作...")
}

// 测试函数
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func (f FlashDisk) deleteData() {
	fmt.Println(f.name, "U盘删除数据...")
}
