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
	// 파일 쓰기 예제
	file, err := os.Create("test_write.txt")
	errCheck1(err)
	defer file.Close()

	// 쓰기 예제 1 - 바이트 슬라이스
	s1 := []byte{115, 111, 109, 101, 115}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", n1)

	// Write Commit(안써도 되긴하나 쓰는 것을 추천, 파일쓰기 이후 중간 저장 느낌)
	file.Sync()

	// 쓰기 예제 2 - 문자열 바로 사용
	s2 := "\nHello Golang!\n File Write Test! - 1\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n", n2)
	file.Sync()

	// 쓰기 예제 3 - 일반 문자열을 바이트로 형변환
	s3 := "Test WriteAt! -2\n"
	// 스페이스바를 50번 누른 위치에서부터 시작
	n3, err := file.WriteAt([]byte(s3), 100) // len(offset) 조절하면서 테스트
	errCheck1(err)

	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n", n3)

	file.Sync()

	// 쓰기 예제 4
	n4, err := file.WriteString("Hello Golang! \n File Write Test! - 3\n")
	errCheck1(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d bytes) \n", n4)

}
