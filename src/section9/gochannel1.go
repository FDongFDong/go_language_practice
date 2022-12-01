// 채널(Channel) 기초(1)

package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1 // 1을 채널로 전송(송신)
}

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2
}
func main() {
	fmt.Println("Main : S ---> ", time.Now())
	// var c chan int
	// c = make(chan int)

	v := make(chan int) // int형 채널 선언
	go work1(v)
	go work2(v)
	<-v
	<-v
	// time.Sleep() 를 사용할 필요가 없다.
	fmt.Println("Main : E ---> ", time.Now())
}
