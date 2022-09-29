package main

import (
	"fmt"
)

func printhello(x int) int {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Hello World")
		}
	}
	return x
}

func main() {
	sum := 0

	// 반복문 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 반복문 2
	for sum < 1000 {
		sum += 1
	}

}
