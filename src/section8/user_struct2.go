package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입
	// 예제 1
	a := cnt(15)

	fmt.Println("ex1: ", a)
	// 예제 2
	var b cnt = 15
	fmt.Println("ex1: ", b)

	// 예제 3
	// testConverT(b) // 에러 int 타입을 넣어야한다.(중요) 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용가능
	testConverT(int(b)) // 형 변환

	// 예제 4
	testConverD(b)
}

func testConverT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
