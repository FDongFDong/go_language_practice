package main

import "fmt"

func main() {
	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// 에러 발생 1
	// 괄호 위치 중요
	// for i := 0; i < 5; i++
	// {
	// 	fmt.Println("ex1 : ", i)
	// }

	// 에러 발생 2
	// 괄호 없으면 안된다.
	// for i := 0; i < 5; i++
	// 	fmt.Println("ex2 : ")

	// 예제 2(무한 루프)
	// for {
	// 	fmt.Println("ex2 : Hello Golang")
	// 	fmt.Println("ex2 : Infinite loop!")
	// }

	// 예제 3(Range용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
}
