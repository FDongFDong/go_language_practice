package main

import "fmt"
import "strconv"

// 함수 선언 위치는 어느곳이든 가능

func helloGolang() {
	fmt.Println("ex1 : Hello Golang!!")
}

func say_one(m string){
	fmt.Println("ex2 : ",m)
}

func sum(x int, y int) int {
	return x+y
}
func main(){
	// 예제 1
	helloGolang()

	// 예제 2
	say_one("Hello World!")

	// 예제 3
	result := sum(5,5)
	fmt.Println("ex3 : ",result)
	fmt.Println("ex3 : ",sum(5,5))
	fmt.Println("ex3 : ",strconv.Itoa(sum(5,5))) // 숫자형을 문자열로 변경
}