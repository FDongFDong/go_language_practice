package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제 1
	kim := Account{"245-901", 1000000, 0.015}
	lee := Account{"245-902", 2000000, 0.025}
	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	CalculateD(kim)
	CalculateP(&lee)
	fmt.Println("ex1 : ", int(kim.balance))
	fmt.Println("ex1 : ", int(lee.balance))
}
