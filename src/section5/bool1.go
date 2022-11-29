// 데이터 타입 : Bool(1)
package main

import "fmt"

func main() {
	// 예제 1
	var b1 bool = true

	var b2 bool = false

	b3 := true

	// b4 := 1

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b3)
	fmt.Println("ex5 : ", b1 && b3) // true
	fmt.Println("ex6 : ", b1 || b2) // true
	fmt.Println("ex7 : ", !b1)      // false

	// 에러
	// 암묵적인 형변환은 이루어지지 않는다.
	// if b4 {
	// 	fmt.Println("ex8 : true")
	// }
}
