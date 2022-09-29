package main

import "fmt"

func main() {
	var num int = 1
	// numPtr에 num 변수의 주소값을 대입한다.
	var numPtr *int = &num

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

}
