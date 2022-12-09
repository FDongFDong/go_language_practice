package main

import "fmt"

func runFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex : ", a[i]) // 에러 발생(인덱스 범위)
	}
}
func main() {

	// 예제
	runFunc()
	// <- 이곳으로 온다.
	fmt.Println("Hello")
}
