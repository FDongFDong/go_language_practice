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
	var num int = 1
	// numPtr에 num 변수의 주소값을 대입한다.
	var numPtr *int = &num

	// [상수]
	const PI = 3.14
	// numPtr이 가지고 있는 num의 주소값이 출력된다.
	fmt.Println(numPtr)
	// numPtr이 가지고 있는 num이 저장하고 있는 값을 출력한다.
	fmt.Println(*numPtr)
	// num의 주소를 출력한다.
	fmt.Println(&num)

	*numPtr = 4
	fmt.Println(*numPtr)

	var numPPtr **int = &numPtr
	fmt.Println(**numPPtr)

	// [알아서 할당]
	var z = 19
	fmt.Println(x, y, PI, z)
}
