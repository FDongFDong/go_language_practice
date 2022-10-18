package main

// [자료구조]
// 문자열
// 배열
// 슬라이스
// 맵
import "fmt"
import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title  string
	Status int
}

func main() {
	// 문자열
	var str string = "abcd"
	var str2 string = "efg"
	fmt.Println(str + str2)
	ExampleTask_marshalJSON()
	ExampleTask_unmarshalJSON()
}

func ExampleTask_marshalJSON() {
	t := Task{ // Task 구조체 선언
		"Laundry", // 초기화
		1,
	}
	b, err := json.Marshal(t) // json 형태의 데이터로 변환 {"Title:"Laundry", "Status":1}가 b에 들어간다. err가 나면 err에 들어간다.
	if err != nil {           // 에러체크
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

	// 배열 선언 및 초기화 1
	var num [5]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	num[3] = 4
	num[4] = 5
	// 배열 선언 및 초기화 2
	num2 := [5]int{1, 2, 3, 4, 5}
	// 배열 선언 및 초기화 3
	num2 := [...]int{1, 2, 3, 4, 5}
func ExampleTask_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2}`)
	t := Task{}                  // 껍데기만 만들어준다.
	err := json.Unmarshal(b, &t) // 껍데기에 넣어준다.
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)

}