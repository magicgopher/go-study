package main

import "fmt"

func main() {

	//
Loop:
	for i := 1; i <= 5; i++ {
		fmt.Printf("i:%d\n", i)
		if i == 3 {
			break Loop
		}
	}

}
