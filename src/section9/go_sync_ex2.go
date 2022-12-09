package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		// 고루틴 추가
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}
	// Add == Done 횟수가 같아야 한다.
	// 끝날때까지 대기
	wg.Wait()
	fmt.Println("WaitGroup End!")
}
