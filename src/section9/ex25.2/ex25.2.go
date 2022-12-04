package main

import (
	"fmt"
	"time"
)

func main() {

	// 채널을 2개 생성 시 빌드하면 2개의 채널 입력값을 받기위해 '9'를 보내고 다음 구문을 진행하기 떄문에 'Never print'가 출력된다.
	// ch := make(chan int, 2)
	ch := make(chan int)
	go square()
	// ch 채널을 가져가는 곳이 없기 때문에 무한히 대기한다.
	ch <- 9
	fmt.Println("Never print")
}
func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
