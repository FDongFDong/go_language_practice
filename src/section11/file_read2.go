package main

import (
	"bufio"
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
	// 파일 읽기
	// csv 파일 읽기 예제

	// 파일 생성
	file, err := os.Open("sample.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// CSV Reader 생성
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// csv 내용 읽기(1)
	row1, err1 := rr.Read() // 1개의 Row 단위로 읽기
	row2, err2 := rr.Read() // 1개의 Row 단위로 읽기
	errCheck1(err1)
	errCheck2(err2)

	fmt.Println("CSV Read Example")
	// fmt.Println(row)
	fmt.Println(row1[0])
	fmt.Println(row2[0])
	fmt.Println("===============================")

	// csv 내용 읽기(2)
	rows, err := rr.ReadAll() // 전체 Row 읽기
	errCheck2(err)
	fmt.Println("CSV ReadAll Example")
	// fmt.Println(rows)
	fmt.Println(rows[5][1])
	fmt.Println("===============================")
	fileInfo, err := file.Stat()
	fmt.Println("파일 정보 출력) : ", fileInfo)
	fmt.Println("파일 이름 : ", fileInfo.Name())
	fmt.Println("파일 크기 : ", fileInfo.Size())
	fmt.Println("파일 수정 시간 : ", fileInfo.ModTime())

	// Row 단위로 CSV 파일 읽기

	for i, row := range rows {
		//fmt.Println(i,row)
		for j := range row {
			fmt.Printf("%s    ", rows[i][j])
		}
		fmt.Println()
	}
}
