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
	// 작업이 끝날때까지 대기
	wg.Wait()

}

func square(wg *sync.WaitGroup, ch chan int) {
	// 계속 가져오려고 하기 때문에 무한대기에 빠진다.
	for n := range ch {
		fmt.Println("Square: ", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
