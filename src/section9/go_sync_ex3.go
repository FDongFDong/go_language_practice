package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	var cnt int64 = 0

	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			wg.Done()
		}(i)
	}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			wg.Done()
		}(i)
	}

	// Add(7000) == Done(7000) 횟수가 같아야 한다.
	// 끝날때까지 대기
	wg.Wait()
	fmt.Println("WaitGroup End! >>>>> ", cnt)
}
