package main

import "fmt"

func main() {
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := true // 예외 발생
	fmt.Println("shortVar1 : ", shortVar1, " shortVar2 : ", shortVar2, " shortVar3 : ", shortVar3)

	// Example
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}
}
