package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 채널 인스턴스 생성
	ch := make(chan int)

	wg.Add(1)
	// 채널자체가 레퍼런스 타입이므로 주소값을 넘길 필요가 없다.
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// ch 채널에서 데이터가 빠져나와 n에 대입한다
	n := <-ch
	time.Sleep(time.Second)
	fmt.Println("Square:", n*n)
	wg.Done()

}
