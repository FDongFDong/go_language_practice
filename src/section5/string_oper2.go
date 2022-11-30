package main

import "fmt"

func main() {
	//예제 2 (비교)
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2)
	fmt.Println("ex2 : ", str1 != str2)
	fmt.Println("ex3 : ", str1 > str2) // 아스키 코드 사전식 비교를 한다. false
	fmt.Println("ex4 : ", str1 < str2) // true
}