package main

import "fmt"

func main() {
	//예제 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1 // sum := sum + 1
	}
	fmt.Println("ex1 : ", sum1)

	//예제 2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
	}
	fmt.Println("ex2 : ", sum2)

	//예제 3
	sum3, i := 0, 0

	for { // while 형태랑 비슷
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	//예제 4
	for i, j := 0, 0; i < 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	//에러 발생
	// 후치 연산은 반환값이 없기 때문에 컴파일 에러
	// for i, j := 0, 0; i < 10; i++, j+=10{
	// }
}
