package main
import "fmt"

func sum(i int, f func(int, int)){
	f(i, 10)
}

func add(a, b int){ // 매개변수가 같은 타입이면 생략가능
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int){
	i = i * 10
}

func multi_reference(i *int){
	*i *= 10 // *i = *i * 10
}

func main(){
	// 함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)

	sum(100, add) // 함수 전달

	// 예제 2
	// call by value
	a := 100
	multi_value(a)
	fmt.Println("ex2 : ",a)

	// 예제 3
	// call by reference
	b := 100
	multi_reference(&b)
	fmt.Println("ex2 : ",b)
}