package main

import ("fmt")

func start(t string) string{
	fmt.Println("start : ",t)
	return t
}

func end(t string) {
	fmt.Println("end : ",t)
}
func a() {
	defer end(start("b")) // defer문은 end함수에만 적용된다. 중첩 함수 주의!
	fmt.Println("in a")
}
func main(){
	// 예제 1
	a()
}