package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 예제 1
	var str1 string = "c:\\go_study\\src\\" // -> c:\go_study\src\
	str2 := `c:\go_study\src\`

	fmt.Println("ex1 :", str1)
	fmt.Println("ex1 :", str2)

	//예제 2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00" // 유니코드

	fmt.Println("ex2 :", str3)
	fmt.Println("ex2 :", str4)
	fmt.Println("ex2 :", str5)

	//예제 3
	//길이(바이트 수)
	fmt.Println("ex3 : ", len(str3)) // 영문은 1바이트
	fmt.Println("ex3 : ", len(str4)) // 한글은 3바이트

	//예제 4
	//길이(실제 길이)
	fmt.Println("ex4 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex4 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex4 : ", len([]rune(str4))) // len으로 실제 길이 구하는 법
}
