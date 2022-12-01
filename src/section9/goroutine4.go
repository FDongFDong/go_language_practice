package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 고루틴(Goroutine)
	// 클로저 사용 예쩨
	// 예제 1
	runtime.GOMAXPROCS(1)
	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시실행 그러나 고루틴으로 실행한 클로저는 가장 나중에 실행(반복문이 종료된 이후 실행됨)
	}

	time.Sleep(5 * time.Second)
}
