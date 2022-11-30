package main

import "fmt"

func main(){
	// 예제 1
	var a *int // 방법 1
	var b *int = new(int) // 방법 2
	

	fmt.Println(a) // nil
	fmt.Println(b) // 주소값

	i := 7 

	a = &i // 주소값 전달
	b = &i 

	*a = 77

	fmt.Println("ex1 : ",a, &i)
	fmt.Println("ex1 : ",&a)
	fmt.Println("ex1 : ",*a) // 역참조

	fmt.Println()

	fmt.Println("ex1 : ",b, &i)
	fmt.Println("ex1 : ",&b)
	fmt.Println("ex1 : ",*b) // 역참조

	var c = &i
	d := &i

	fmt.Println("ex1 : ",c, &i)
	fmt.Println("ex1 : ",&c)
	fmt.Println("ex1 : ",*c) // 역참조

	fmt.Println()

	fmt.Println("ex1 : ",d, &i)
	fmt.Println("ex1 : ",&d)
	fmt.Println("ex1 : ",*d) // 역참조

}