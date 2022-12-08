package main

import (
	"errors"
	"fmt"
)

// 죽는 방법
func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

// 사는 방법
// func divide(a, b int) error {
// 	if b == 0 {
// 		return errors.New("b는 0일 수 없습니다.")
// 	}
// 	fmt.Printf("%d / %d = %d\n", a, b, a/b)
// 	return nil
// }

func main() {
	divide(9, 3)
	divide(9, 0)
}

// 9 / 3 = 3
// panic: b는 0일 수 없습니다.

// goroutine 1 [running]:
// main.divide(0x9?, 0x3?)
//         /Users/jhka/Documents/fdongfdong/STUDY/go/src/section10/ex23.5/ex23.5.go:7 +0x105
// main.main()
//         /Users/jhka/Documents/fdongfdong/STUDY/go/src/section10/ex23.5/ex23.5.go:14 +0x31
