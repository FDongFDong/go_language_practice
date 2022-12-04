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
	// 초당 한번씩 신호를 보낸다
	tick := time.Tick(time.Second)
	// 10초 뒤 한번 신호를 보낸다.
	terminate := time.After(12 * time.Second)
	for {
		select {
		case n1 := <-tick:
			fmt.Println("n1 tick! : ", n1)
		case n2 := <-terminate:
			fmt.Println("n2 terminated : ", n2)
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square : ", n*n)
			time.Sleep(1 * time.Second)
		}
	}
}
