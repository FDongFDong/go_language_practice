package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 구조체 선언(공유 데이터 예시)
type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	// 공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	// 공유 데이터 수정 후 뮤텍스 보호 해제
	c.mutex.Unlock()
}

// 읽기만 할것이므로 굳이 포인터로 사용안해도되지만 가독성, 유지보수를 위해 그냥 포인터로 선언
func (c *count) result() {
	fmt.Println(c.num)
}
func main() {
	// 시스템 전체 cpu 사용 -> 경쟁 상태가 더 심해짐
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}

	done := make(chan bool)
	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			// 다른 cpu에게 양보
			runtime.Gosched()
		}()
	}
	for i := 1; i <= 10000; i++ {
		<-done
		c.result()
	}
}
