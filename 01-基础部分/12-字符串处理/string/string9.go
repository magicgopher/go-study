package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 定义字符串
	s1 := "hello world!"
	sub1 := s1[0:5]
	sub2 := s1[6:12]
	// 打印sub1和sub2
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println("=======================")
	// 打印原字符串的子字符串[sub1 sub2]的地址
	// sub1 是从 s1 中截取的前5个字符 "hello"，所以 sub1 的底层数据地址与 s1 的起始地址相同。
	fmt.Printf("s1地址:%p\n", unsafe.Pointer(&s1))
	fmt.Printf("sub1地址:%p\n", unsafe.Pointer(&sub1))
	fmt.Printf("sub2地址:%p\n", unsafe.Pointer(&sub2))
	fmt.Println("=======================")
	// 使用 reflect 包查看底层数据的地址
	headerS1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	headerSub1 := (*reflect.StringHeader)(unsafe.Pointer(&sub1))
	headerSub2 := (*reflect.StringHeader)(unsafe.Pointer(&sub2))
	// 打印底层地址
	// sub2 是从 s1 中截取的字符 "world!"，它从 s1 的第6个字符开始（索引从0开始）
	// 所以 sub2 的底层数据地址是 s1 的起始地址加上6个字符的偏移量，即 地址 + 6
	fmt.Printf("s1底层地址:%p\n", unsafe.Pointer(headerS1.Data))
	fmt.Printf("sub1底层地址:%p\n", unsafe.Pointer(headerSub1.Data))
	fmt.Printf("sub2底层地址:%p\n", unsafe.Pointer(headerSub2.Data))
}
