package main

import "fmt"

func rptc(n *int){
	*n = 77
}

func vptc(n int){
	n = 77
}

func main(){
	var a int = 10
	var b int = 10

	fmt.Println("ex1 : ",a)
	fmt.Println("ex1 : ",b)
	fmt.Println()

	rptc(&a)
	vptc(b)
	// vptc(&b) // 에러 발생

	fmt.Println("ex2 : ",a)
	fmt.Println("ex2 : ",b)
	


}