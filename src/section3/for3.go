package main

import "fmt"

func main() {
	// 예제 1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1: ", i, j)
		}
	}
	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 : ", i)
	}

Loop2:
	// 에러 발생
	// Loop 레이블 아래에는 조건문 또는 반복문이 와야한다. fmt.Println()이 오면 안된다.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3 : ", i, j)
		}
	}
}
