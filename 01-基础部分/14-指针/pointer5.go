package main

import "fmt"

func main() {
	x, y := 10, 20
	fmt.Println("交换前 x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("交换后 x =", x, "y =", y)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
