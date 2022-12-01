//고루틴(Goroutine) 기초(1)

package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe3 func end", time.Now())
}

func main() {

	exe1() // 가장 먼저 실행(일반적인 실행 흐름)

	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
