package main

import (
	"fmt"
)

func receiveOnly(cnt int) <-chan int {
	sum := 1
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}

	}()
}
func main() {
	// 채널

	// 예제 1

	c := receiveOnly(100) // 채널 반환
	output := total(c)    // 채널 전달 후 반환

	fmt.Println("ex1 : ", <-output)
}
