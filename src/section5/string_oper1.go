package main

import "fmt"

func main() {
	//예제 1(추출)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0]) // Go 71
	fmt.Println("ex1 : ", str2[3:], str2[0]) // ld 87
	fmt.Println("ex1 : ",str2[:4]) // Worl
	fmt.Println("ex1 : ",str1[1:3]) // ol


}