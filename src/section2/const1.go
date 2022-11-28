package main

import "fmt"

func main() {
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight() // 항상 같은 값을 보장할 수 없으므로 Error
	const e = 35.6
	const f = false
	/*
		에러 발생
		cosnt g string
		g = "Test"
	*/
	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
}
