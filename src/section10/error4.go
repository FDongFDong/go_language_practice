package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golan : ", n)
		return s, nil
	}
	return "", errors.New("0을 입력했습니다. 에러 발생!")

}

// Errorf를 이용한 에러 처리
func main() {
	a, err := notZero(1)
	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}
	fmt.Println("ex1 : ", a)

	b, err := notZero(0)
	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}

	//Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("ex2 : ", b)
	fmt.Println("End Error Handling!")

}
