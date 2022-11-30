package main

import "fmt"

func main(){
	// 맵 값 변경 및 삭제
	map1 := map[string]string {
		"daum": "http://daum.net",
		"naver": "http://naver.com",
		"google": "http://google.com",
		"home1": "http://test1.com",
	}


	//예제 1
	fmt.Println("ex1 : ", map1)
	map1["home2"] = "http://test2.com" // 추가
	fmt.Println("ex1 : ",map1)

	map1["home2"] = "http://test2-2.com" // 수정
	fmt.Println("ex1 : ",map1)

	fmt.Println()

	//예제2 (삭제)
	delete(map1, "home2")
	fmt.Println("ex1 : ", map1)

	delete(map1, "google")
	fmt.Println("ex1 : ", map1)
	
}