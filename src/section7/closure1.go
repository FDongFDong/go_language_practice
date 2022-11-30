package main

import "fmt"

func main(){
	// 예제 1
	multiply := func(x,y int) int { // 익명 함수
		return x * y
	}
	r1 := multiply(5,10)
	fmt.Println("ex1 : ", r1)

	// 예제 2

	m, n := 5,10 // 이때 변수가 캡쳐된다.
	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c // 지역 변수가 소멸되지 않는다. (함수 호출 시 마다 사용 가능)
	}
	r2 := sum(10)
	fmt.Println("ex2 : ",r2)
}