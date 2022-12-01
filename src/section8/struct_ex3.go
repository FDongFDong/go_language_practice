package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리: 구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{"245-901", 1000000, 0.015}
	lee := Account{"245-902", 2000000, 0.025}
	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	fmt.Println()
	lee.CalculateD(100000)
	kim.CalculateF(100000)

	fmt.Println("ex1 : ", int(lee.balance))
	fmt.Println("ex1 : ", int(kim.balance))
}
