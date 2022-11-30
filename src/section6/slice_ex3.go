package main
import "fmt"

func main(){
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// 반드시 make로 공간을 할당 후 복사해야 한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	slice1 := []int{1,2,3,4,5,6,7,8,9,10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1) // slice의 공간만큼만 복사된다.
	copy(slice3, slice1) // 복사가 안된다.

	fmt.Println("ex1 : ",slice2)
	fmt.Println("ex1 : ",slice3)

	// 예제 2
	a := []int{1,2,3,4,5}
	b := make([]int, 5)

	copy(b,a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ",a)
	fmt.Println("ex2 : ",b)
	fmt.Println()

	// 예제 3
	c := [5]int{1,2,3,4,5}
	d := c[0:3]

	d[1] = 7

	fmt.Println("ex3 : ",c)
	fmt.Println("ex3 : ",d)
	fmt.Println()

	// 예제 4
	e := []int{1,2,3,4,5,6,7,8,9,10}
	f := e[0:5:7]	 // 용량까지 지정할 수 있음

	fmt.Println("ex4 : ",len(f), cap(f))
	fmt.Println("ex4 : ",f)



}