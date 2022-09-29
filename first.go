package main

// [자료구조]
// 문자열
// 배열
// 슬라이스
// 맵
import "fmt"

func main() {
	// 문자열
	var str string = "abcd"
	var str2 string = "efg"
	fmt.Println(str + str2)

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
