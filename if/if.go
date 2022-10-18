package main

import "fmt"

// func 이름(입력값 정의) 출력값 정의 {
//  함수 내용
//}

// call by value
func sum(x int, y int) int {
	return x + y
}

// call by reference
func set(x *int) int {
	*x = *x + 3
	return *x
}

func add(x int, y int) (int, int) {
	var c int = x + y
	return c, 9
import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func main() {
	var c int = sum(1, 2)
	var x = 5
	fmt.Println(set(&x))
	fmt.Println(c)

	a, b := add(5, 3)
	e, f := 2, 5
	fmt.Println(a, b)
	fmt.Println(e, f)
	// 동적타입은 아니다. 재할당시 스트링만 가능
	g := "str"
	fmt.Println(g)

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}