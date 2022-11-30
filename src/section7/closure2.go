package main

import "fmt"

func main(){
	// 예제 1
	cnt := increaseCnt()
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())

}

func increaseCnt() func() int {
	n := 0 // 지역 변수(캡쳐됨)
	return func() int {
		n +=1 
		return n
	}
}