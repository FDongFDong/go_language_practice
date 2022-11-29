package main

import "fmt"

func main() {
	//코드 서식 지정
	//한 사람이 코딩 한 것 같은 일관성 유지
	//코드 스타일 유지
	//gofmt -h : 사용법
	//gofmt -w abc.go: 원본파일에 반영

	//예제 1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex1 : ", i)
	}
}
