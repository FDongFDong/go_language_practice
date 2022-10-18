// [package 이름을 정하는 부분]
// 내가 원하는 이름을 정할 수 있다.
// 파일이름은 그냥 파일이름일 뿐이고 package에 적혀있는 이름이 실제 이름이다.
// second.go 파일에서 package 이름을 main으로 두면 같은 파일이 된다.
// * 패키지에서 main은 특별한 패키지이다.
// [go 규칙]
// - 모든 프로그램은 main패키지 부터 시작해야한다.
// - 모든 프로그램은 하나 이상의 main 패키지가 있어야 한다.
package main

// ----------------------
// [외부 패키지 불러오는 곳]
import (
	"fmt"
)

// ----------------------
// [source code]
// main 함수 -
func main() {
	// fmt 패키지에 있는 Println() 함수를 호출한다.
	fmt.Println("Hello My Class !")
}
	// [변수]
	// int형 정수 x에 12를 대입한다.
	// 임의 메모리 위치에 x라는 이름으로 가리키고 그곳에 12를 저장한다.
	var x int = 12
	// x라는 메모리위치에 값을 가져온다.
	var y float64 = 3.21

	// [상수]
	const PI = 3.14

// 프로그램 실행 전 컴파일 작업 필요
// go mod init "test"
	// [알아서 할당]
	var z = 19
	fmt.Println(x, y, PI, z)
}