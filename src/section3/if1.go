package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF문 : 반드시 Boolean형으로 검사
	//   - 0, 1 사용불가
	// 소괄호를 사용하지 않는다.
	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
	}
	fmt.Println("15이상")
	// 예제 2
	if b >= 25 {
		fmt.Println("25이상")
	}
	// 에러 발생 1
	// 괄호는 if 문 마지막에 둬야한다. 코드 컨벤션
	// if b >= 25
	// {
	// }

	// 에러 발생 2
	// 괄호 없이 사용 불가능
	// if b >= 25
	// 	fmt.Println("25이상")

	// 에러 발생 3
	// 1과 0을 true, false로 사용할 수 없다.
	// if c:= 1; c{
	// 	fmt.Println("True")
	// }

	// 예제 3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}
	// c는 에러
	// c += 20
}
