package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// return fmt.Errorf("...")
		// return errors.New("...")
		return PasswordError{len(password), 8}
	}
	// ...
	return nil
}
func main() {
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		// 인터페이스 객체로 변환될 수 있으면
		if errorInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errorInfo, errorInfo.Len, errorInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
