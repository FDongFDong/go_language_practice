package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // 임원도 직원, is a 관계
	specialBonus float64
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메서드를 재 사용하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메서드 상속 활용을 위한 패턴

	// 예제 1
	// 직원
	ep1 := Employee{"kim", 200000, 150000}
	ep2 := Employee{"park", 150000, 200000}

	// 임원
	ex := Executives{
		Employee{"lee", 500000, 100000},
		100000,
	}

	fmt.Println("ex1 :", int(ep1.Calculate()))
	fmt.Println("ex1 :", int(ep2.Calculate()))

	// Employee 구조체를 통해서 메서드 호출
	fmt.Println("ex : ", int(ex.Calculate()+ex.specialBonus))
}
