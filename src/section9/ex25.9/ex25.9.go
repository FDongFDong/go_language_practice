package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// number라는 키값에 9라는 데이터를 넣어 지정한다.
	// 작업을 지시할 때 특정 데이터를 통해 지시할 수 있다.
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		// 리턴 값은 빈 인터페이스이기 떄문에 타입 변환을 해줘야한다.
		n := v.(int)
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}
