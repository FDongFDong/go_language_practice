package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	// 한줄씩, 한단어씩 읽어오기 편함 / strings.NewReader: io.Reader 객체를 만들어서 반환해준다.
	scanner := bufio.NewScanner(strings.NewReader(str))
	// 한 단어씩
	scanner.Split(bufio.ScanWords)
	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자, 읽은 글자수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	// 한단어를 읽어옴
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	// 문자열을 숫자로 바꾼다. "24" -> 24
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, erro:%w", word, err)
	}
	return number, len(word), nil
}
func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// err로 감싸진 것중에 numError이 있을 경우 꺼내와서 읽어준다.
		if errors.As(err, &numError) {
			fmt.Println("numberError", numError)
			fmt.Println("numberError", numError.Func)
		}
	}

}
func main() {
	readEq("123 3")
	readEq("123 abc")
}
