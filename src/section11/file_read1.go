package main

import (
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
	// Open :  기존 파일 열기
	// Clse : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요
	// 탐색 Seek 중요

	// 파일 읽기 예제
	// 파일 열기
	file, err := os.Open("sample.txt")
	errCheck1(err)
	defer file.Close()

	// 읽기 예제 1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득

	errCheck2(err)
	// 가져올 내용만큼 슬라이스 선언하고 읽은 내용을 담는다.
	fd1 := make([]byte, fileInfo.Size())

	ct1, err := file.Read(fd1)
	errCheck1(err)
	fmt.Println("파일 정보 출력(1) : ", fileInfo, "\n")
	fmt.Println("파일 이름(2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)
	// fmt.Println(fd1) 바이트이기 떄문에 형변환 필요
	fmt.Println(string(fd1))
	fmt.Println("=========================================")

	// 읽기 예제2 (탐색 : Seek(Offset))
	o1, err := file.Seek(20, 0) // 20만큼 커서가 이동한 다음 / 0 : 처음위치, 1: 현재위치, 2: 마지막위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")
	fmt.Println("=========================================")

	// 읽기 예제3
	o2, err := file.Seek(0, 0)
	errCheck1(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.
	errCheck1(err)
	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
	fmt.Println("=========================================")

}
