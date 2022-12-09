package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	var cnt int64 = 0

	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			// cnt += 1
			// 1씩 더한다
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			// cnt -= 1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	// Add(7000) == Done(7000) 횟수가 같아야 한다.
	// 끝날때까지 대기
	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	// 방법 1
	fmt.Println("WaitGroup End! >>>>> ", cnt)
	// 방법 2 - 추천! 계산이 끝난
	fmt.Println("WaitGroup End! >>>>> ", finalCnt)
}
