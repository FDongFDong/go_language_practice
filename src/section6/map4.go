package main

import "fmt"

func main(){
	// 맵 조회 할 경우 주의할 점

	// 예제 1
	map1 := map[string]int { // int : 0, string : "", float : 0.0
		"apple" : 15,
		"banana": 115,
		"orange": 1115,
		"lemon": 0,
	}

	value1 := map1["lemon"] 
	value2 := map1["kiwi"] // key가 없으면 해당 타입의 초기값이 나온다
	value3, ok := map1["kiwi"] // 두번째는 키가 존재하는지 true, false가 나온다

	fmt.Println("ex1 : ",value1)
	fmt.Println("ex2 : ",value2)
	fmt.Println("ex3 : ",value3, ok) 

	// 예제 2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	}else{
		fmt.Println("ex2 : kiwi is not exist!")
	}
	
	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	}else{
		fmt.Println("ex2 : banana is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok { // false면
		fmt.Println("ex2 : kiwi is not exist!")
	}
}