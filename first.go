package main

import (
	"fmt"
)

func main() {
	// [변수]
	// int형 정수 x에 12를 대입한다.
	// 임의 메모리 위치에 x라는 이름으로 가리키고 그곳에 12를 저장한다.
	var x int = 12
	// x라는 메모리위치에 값을 가져온다.
	var y float64 = 3.21

	// [상수]
	const PI = 3.14

	// [알아서 할당]
	var z = 19
	fmt.Println(x, y, PI, z)
}
