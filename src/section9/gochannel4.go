package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널(Channel)
	// 예제 1(비동기: 버퍼 사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기, 가득 차있으면 작동

	runtime.GOMAXPROCS(4)
	ch := make(chan bool, 4) // 버퍼의 용량 설정 가능, 버퍼가 가득 찰때까지 보내고 대기
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)

		}
	}()
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
