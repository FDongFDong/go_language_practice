package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xLsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// csv Writer 생성
	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewWriter(file)) <- 권장

	// csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "4.2"})
	wr.Write([]string{"Park", "4.3"})
	wr.Write([]string{"Cho", "4.1"})
	wr.Write([]string{"Hong", "4.4"})

	wr.Flush() // 버퍼 -> 파일로 쓰기
	// 파일로 바로 쓰고 열고하는 것은 io가 많이 유발되기 때문에 버퍼에 한번 딱 쓰고 버퍼가 차면 버퍼를 비우는 식으로 하는 것이 좋다

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일 명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한: ", fi.Mode())

}
