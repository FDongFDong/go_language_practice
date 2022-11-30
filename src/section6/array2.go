package main

import "fmt"

func main(){
	// 배열 순회

	// 예제 1
	arr1 := [5]int{1,10,100,1000,10000}

	//len 길이 반복
	for i := 0; i<len(arr1) ; i++{
		fmt.Println("ex1 : ",arr1[i])
	}
	fmt.Println()

	// 예제 2
	arr2 := [5]int{7,77,777,7777,77777}
	//range
	for i, v:= range arr2 {
		fmt.Println("ex2 : ",i,v)
	}

	// 인덱스 생략1
	for _, v:= range arr2 {
		fmt.Println("ex2 : ", v)
	}

	// 값 생략2
	fmt.Println()
	for v:= range arr2 {
			fmt.Println("ex2 : ", v)
	}
}
