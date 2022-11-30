//사용자 정의 타입(1)

package main
import "fmt"

type Car struct {
	name string
	color string
	price int64
	tax int64
}


// 일반 메서드
func Price(c Car) int64 {
	return c.price + c.tax
}
// 구조체 <-> 메서드 바인딩
func (c Car) Price() int64{
	return c.price + c.tax
}

func main() {
	// Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	// 객체지향 -> 클래서(속성 : 멤버변수, 기능(상태 : 메서드)) : 코드의 재사용성, 코드의 관리, 신뢰성이 높은 프로그래밍
	// Go는 전형적인 객제지향의 전형적인 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스형태의 코딩 가능
	// 객체지향의 기본 개념 -> Go에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
	// 상태, 메서드 분리해서 정의(결합성 없음)
	// 사용자 정의 타입: 구조체, 인터페이스, 기본 타입(int, float, string...), 함수
	// 구조체와 메서드 연결을 통해서 타 언어의 클래스 형식처럼 사용가능(객체 지향)
	
	// 예제 1
	bmw := Car{name:"520d", price:500000000, color: "white", tax: 50000000}
	benz := Car{name:"220d", price:600000000, color: "white", tax: 60000000}

	fmt.Println("ex1 : ",bmw, &bmw)
	fmt.Println("ex1 : ",benz, &benz)
	// fmt.Printf("ex1 :  %p",&bmw, "\n")
	// fmt.Printf("ex1 :  %p",&benz, "\n")

	// fmt.Println("ex2 : ", Price(bmw))
	// fmt.Println("ex2 : ", Price(benz))

	fmt.Println("ex3 : ", bmw.Price())
	fmt.Println("ex3 : ", benz.Price())

	fmt.Print("ex4 : ",&bmw == &benz)


}