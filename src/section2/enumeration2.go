package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
	)

	const (
		Jan = iota + 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6
	)
	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)

	fmt.Println(A, B, C)
}
