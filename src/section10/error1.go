package main

import (
	"fmt"
	"log"
	"os"
)

// Fatal() : 프로그램 종료
func main() {
	// 일부러 예외 발생
	f, err := os.Open("unnamedfile")
	if err != nil {
		log.Fatal(err.Error())
		// log.Fatal(err) 위와 같은 출력
	}
	fmt.Println("ex1 : ", f.Name())

}
