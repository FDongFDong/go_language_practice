package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func SumAtoB(a,b int) {
	sum := 0
	for i:= a ; i<=b ;i++{
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n",a,b,sum)
	wg.Done() // 하나의 작업이 끝났음을 알려주는 함수
}

func main(){
	wg.Add(10) // 10개의 작업을 기다린다.
	for i := 0 ;i<10; i++{ // 10개의 goroutine가 생긴다.
		go SumAtoB(1, 1000000000)
	}
	wg.Wait() // 10개가 끝날 때까지 프로그램을 종료하지 않는다.
}