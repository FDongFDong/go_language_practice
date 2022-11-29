# go_language_practice

- [go\_language\_practice](#go_language_practice)
  - [Go 언어 장점 및 특징](#go-언어-장점-및-특징)
  - [변수 및 상수](#변수-및-상수)
  - [제어문 및 반복문](#제어문-및-반복문)
    - [if](#if)
    - [switch](#switch)
    - [for](#for)
  - [조건문](#조건문)
  - [함수](#함수)
  - [조건문](#조건문-1)
  - [반복문](#반복문)
  - [자료구조 - 문자열 배열](#자료구조---문자열-배열)
  - [marshal unMarshal](#marshal-unmarshal)
  - [Ethereum Core Source를 이용한 Explorer 개발](#ethereum-core-source를-이용한-explorer-개발)

## Go 언어 장점 및 특징

- 간결한 문법 및 단순함
- 병행 프로그래밍 지원
- 정적 타입 및 동적 실행
- 간편한 협업 지원
- 컴파일 및 실행속도 빠름
- 제네릭 및 예외 처리 미지원
- 컨벤션 통일

___

## 변수 및 상수

> [variable1.go]()
>
> [variable2.go]()
>
> [variable3.go]()
>
> [const1.go]()
>
> [const2.go]()
>
> [enumeration1.go]()
>
> [enumeration2.go]()

- 기본 초기화
  - 정수 타입: 0,
  - 실수(소수점) : 0.0,
  - 문자열 : "",
  - Boolean: true, false
  - 변수명 : 숫자 첫글자x, 대소문자 구분o, 문자, 숫자 밑줄, 특수기호 사용 가능
  - 변수 및 상수 : 함수 내외 사용 가능
- 여러개 선언 하기
- 짧은 선언
  - 함수 안에서만 사용 가능(전역 x)
  - 선언 후 할당 시 예외 발생
  - 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있다.
- 상수
  - const 예약어를 통해 사용 초기화
  - 한 번 선언 후 값 변경 금지
  - 고정 된 값 관리용
- 열거형
  - 상수를 사용하는 일정한 규칙에 따라 숫자를 계산하거나 증가시키는 묶음

## 제어문 및 반복문
### if 

> [if1.go]()
> [if2.go]()
> [if3.go]()

- 괄호는 if문 뒤에 둬야한다.

### switch

> [switch1.go]()
> [switch2.go]()
> [switch3.go]()

- Switch 키위드 뒤 표현식(Expression) 생략 가능
- case 뒤 표현식(Expression) 사용 가능
- 자동 break 때문에 fallthrough 존재
- Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

### for
> [for1.go]()
> [for2.go]()
> [for3.go]()

- while, do while문이 없다.
- Go에서 유일하게 제공되는 반복문
- 다양한 사용법 숙지 필요


___

## 조건문
> [pointer.go](https://github.com/FDongFDong/go_language_practice/blob/main/pointer/pointer.go)
___

## 함수

> [function.go](https://github.com/FDongFDong/go_language_practice/blob/main/function/func.go)
___

## 조건문

> [if.go](https://github.com/FDongFDong/go_language_practice/blob/main/if/if.go)
___

## 반복문

> [forloop.go](https://github.com/FDongFDong/go_language_practice/blob/main/forloop/for.go)
___

## 자료구조 - 문자열 배열

> [structure.go](https://github.com/FDongFDong/go_language_practice/blob/main/structure/data_structure.go)
___

## marshal unMarshal

> [json](https://github.com/FDongFDong/go_language_practice/blob/main/json/json.go)
___

## Ethereum Core Source를 이용한 Explorer 개발

- Block event 데이터 처리하기
[go-ethereum/ethclient](https://github.com/ethereum/go-ethereum/tree/master/ethclient)
[go-ethereum/core/types](https://github.com/ethereum/go-ethereum/tree/master/core/types)

> [BlockListner.go](https://github.com/FDongFDong/go_language_practice/tree/main/explorer)

- Block 정보 가져오기
- Transaction 정보 가져오기
- ERC20 토큰 정보 가져오기

<img width="1765" alt="image" src="https://user-images.githubusercontent.com/20445415/201647346-b1a44e7d-a701-49a4-ba6a-e1b2dff84a40.png">
