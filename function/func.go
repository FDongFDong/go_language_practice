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
}

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

}