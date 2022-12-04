package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main(){
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	for i := 0; i< 10; i++{ // 10개의 고루틴
		go func(){ 
			for { // 입금하고 출금하는 함수를 무한히 반복한다.
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 1000원을 입금하고 출금하는 함수
func DepositAndWithdraw(account *Account){
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000	 // 10개의 고루틴이 account.Balance에 접근한다.
	//account.Balance = account.Balance + 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
	//account.Balance = account.Balance - 1000
}

// 하나의 메모리 자원에 여러개의 고루틴이 접근해서 생기는 에러