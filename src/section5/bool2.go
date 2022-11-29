// 데티어 타입 : bool(2)
package main

import "fmt"

func main() {
	// 예제 1(논리 연산자)
	fmt.Println("ex1 : ", true && true)   // true
	fmt.Println("ex1 : ", true && false)  // false
	fmt.Println("ex1 : ", false && false) // false
	fmt.Println("ex1 : ", true || true)   // true
	fmt.Println("ex1 : ", true || false)  // true
	fmt.Println("ex1 : ", false || false) // false
	fmt.Println("ex1 : ", !true)          // false
	fmt.Println("ex1 : ", !false)         // true

	// 예제 2(비교 연산자)
	num1 := 15
	num2 := 37

	fmt.Println("ex2 : ", num1 < num2)
	fmt.Println("ex2 : ", num1 > num2)
	fmt.Println("ex2 : ", num1 >= num2)
	fmt.Println("ex2 : ", num1 <= num2)
	fmt.Println("ex2 : ", num1 == num2)
	fmt.Println("ex2 : ", num1 != num2)
}
