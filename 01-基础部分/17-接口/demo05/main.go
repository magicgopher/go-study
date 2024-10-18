package main

import (
	"fmt"
	"math"
)

func main() {
	// 创建一个正方形实例
	s1 := &Square{side: 5}

	// 将 s1 赋值给 Shaper 接口
	var shaper Shaper
	shaper = s1

	// 判断 shaper 中的具体类型
	if t, ok := shaper.(*Square); ok {
		fmt.Printf("The type of shaper is: %T\n", t)
		fmt.Println("Area of square:", t.Area())
	}
	if u, ok := shaper.(*Circle); ok {
		fmt.Printf("The type of shaper is: %T\n", u)
		fmt.Println("Area of circle:", u.Area())
	} else {
		fmt.Println("shaper does not contain a variable of type Circle")
	}
}

// Square 表示正方形
type Square struct {
	side float32 // 边长
}

// Circle 表示圆形
type Circle struct {
	radius float32 // 半径
}

// Shaper 表示形状的接口
type Shaper interface {
	Area() float32 // 计算面积的方法
}

// Area 计算正方形的面积
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// Area 计算圆形的面积
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
