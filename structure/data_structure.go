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
// [자료구조]
// 문자열
// 배열
// 슬라이스
// 맵
import "fmt"

func main() {
	sum := 0

	// 반복문 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 문자열
	var str string = "abcd"
	var str2 string = "efg"
	fmt.Println(str + str2)

	// 반복문 2
	for sum < 1000 {
		sum += 1
	}
	// 배열 선언 및 초기화 1
	var num [5]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	num[3] = 4
	num[4] = 5
	// 배열 선언 및 초기화 2
	num2 := [5]int{1, 2, 3, 4, 5}
	// 배열 선언 및 초기화 3
	num2 := [...]int{1, 2, 3, 4, 5}

}