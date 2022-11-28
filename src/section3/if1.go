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
	// if b >= 25 
	//{
	// }
  // 에러 발생 2
	if b >= 25
	fmt.Println("25이상")
}
