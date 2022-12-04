package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// 데이터를 모두 보내면 채널을 닫아준다.
	close(ch)
	// 작업이 끝날때까지 대기
	wg.Wait()

}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square: ", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
