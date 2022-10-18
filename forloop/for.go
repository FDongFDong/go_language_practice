package main

import (
	"fmt"
)

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
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
