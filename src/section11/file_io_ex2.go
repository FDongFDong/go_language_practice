package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 없으면 만들고, 있으면 읽고 써라
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	// 에러 체크
	errCheck(err)

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) // Writer 반환(버퍼 사용)
	wt.WriteString("Hello Golang!\n File Write Test1!\n")
	wt.Write([]byte("Hello Golang!\n File Write Test2!\n"))

	// 버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size (%d byte)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d byte)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d byte)\n", wt.Size())

	// 버퍼 비우고 반영(버퍼에 내용을 디스크에 기록)
	wt.Flush()
	fmt.Println("쓰기 작업 완료")
	fmt.Println("=============================")

	rt := bufio.NewReader(file) // Reader 반환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) // 읽기(ReadLine, ReadByte, ReadBytes 등)
	// re.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("=============================")
	fmt.Println(string(b))
	fmt.Println("=============================")
	defer file.Close()

}
