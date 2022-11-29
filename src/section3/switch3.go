package main

import "fmt"

func main() {
	// 예제 1
	a := 30 / 15 // 2
	switch a {
	case 2, 4, 6: // i가 2, 4, 6일 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // i가 1, 3, 5일 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	// 예제 2
	// fallthrough
	// 조건에 맞는 문장 아래에 반드시 실행되는 문장이 있다면 fallthrough 사용
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")

	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")

	case "ruby":
		fmt.Println("ruby!")

	case "c":
		fmt.Println("c!")
	}
}
