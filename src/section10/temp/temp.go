package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	// 파일 핸들러와 에러 발생시 에러 리턴
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// reader 인스턴스를 리턴한다. NewReader은 read 함수만 받는 인터페이스, file역시 read 함수를 가지고 있기 떄문에 가능하다
	rd := bufio.NewReader(file)
	// 데이터를 읽어온다, 개행문자까지(한줄만 읽어오겠다). 스트링으로 받아온다. 에러로는 EOF가 나온다(다 읽었다라는 의미. 크게 중요치 않아서 무시)
	line, _ := rd.ReadString('\n')
	// rd.ReadLine() 다른점은 바이트 배열로 나온다.
	// 읽었으면 읽은 한줄하고 에러가 없음을 알려주는 nil 반환
	return line, nil
}

func WriteFile(filename string, line string) error {
	// err는 권한문제나 디스크 에러가없을 떄 등
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Fprintln() // 첫번째 인자로 어느곳에 쓸것인지 정해준다
	_, err = fmt.Fprint(file, line)
	// 에러가 없으면 nil 반환
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		// 파일을 읽는데 실패했으면 파일을 써라
		err = WriteFile(filename, "This is WriteFile")
		// 만약 쓰는 것도 실패했으면
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line)
}
