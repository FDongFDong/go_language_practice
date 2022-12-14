package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)
	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Wating End : ", n)
			mutex.Unlock()
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-c
		// fmt.Println("received : ", <-c)
	}

	// 하나씩 꺠우기
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Goroution(Signal) : ", i)
	// 	condition.Signal() // 한 개 씩 깨운다.(모든 고루틴 생성 후)
	// 	mutex.Unlock()
	// }
	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
