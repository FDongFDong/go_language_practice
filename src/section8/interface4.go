package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메서드 구현

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

func (d Cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}
func act(animal interface{ run() }) { // 익명 인터페이스 선언(타입 정의X)
	animal.run()
}

func main() {

	// 익명 인터페이스 사용 예제(즉시 선언 후 사용)

	// 예제 1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}
	// 개 행동 실행
	act(dog)
	// 고양이 행동 실행
	act(cat)

}
