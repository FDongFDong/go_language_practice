package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// 시스템 전체 cpu 사용
	data := 0

	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMutex)

	// 쓰는 작업
	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			// 쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()
	// 쓰는 도중에 읽기
	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 해제
			mutex.RUnlock()
		}
	}()
	// 쓰는 도중에 읽기
	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 해제
			mutex.RUnlock()
		}
	}()
	time.Sleep(10 * time.Second)
}
