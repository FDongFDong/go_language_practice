# go_language_practice

- [go\_language\_practice](#go_language_practice)
  - [Go 언어 장점 및 특징](#go-언어-장점-및-특징)
  - [변수 및 상수](#변수-및-상수)
  - [제어문 및 반복문](#제어문-및-반복문)
    - [if](#if)
    - [switch](#switch)
    - [for](#for)
  - [패키지](#패키지)
  - [접근제어 및 Alias](#접근제어-및-alias)
  - [초기화 메소드(init)](#초기화-메소드init)
  - [데이터 타입](#데이터-타입)
    - [Bool](#bool)
    - [숫자형 기초](#숫자형-기초)
    - [숫자형 연산](#숫자형-연산)
    - [문자열 기초](#문자열-기초)
    - [문자열 연산](#문자열-연산)
  - [자료구조](#자료구조)
    - [배열](#배열)
  - [조건문](#조건문)
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

> [variable1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable1.go)
>
> [variable2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable2.go)
>
> [variable3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/variable3.go)
>
> [const1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/const1.go)
>
> [const2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/const2.go)
>
> [enumeration1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration1.go)
>
> [enumeration2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration2.go)
>
> [enumeration3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section2/enumeration3.go)

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

> [if1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if1.go)
>
> [if2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if2.go)
>
> [if3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/if3.go)

- 괄호는 if문 뒤에 둬야한다.

### switch

> [switch1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch1.go)
> 
> [switch2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch2.go)
> 
> [switch3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/switch3.go)

- Switch 키위드 뒤 표현식(Expression) 생략 가능
- case 뒤 표현식(Expression) 사용 가능
- 자동 break 때문에 fallthrough 존재
- Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

### for
> [for1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for1.go)
>
> [for2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for2.go)
>
> [for3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section3/for3.go)

- while, do while문이 없다.
- Go에서 유일하게 제공되는 반복문
- 다양한 사용법 숙지 필요


___

## 패키지

> [package1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/package1.go)
>
> [package2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/package2.go)

- 패키지 : 코드 구조화(모듈화) 및 재사용
- 응집도, 결합도 
- Go : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
- 패키지 이름 = 디렉토리 이름
- 같은 패키지 내 -> 소스 파일들은 디렉토리명을 패키지 명으로 사용한다.
- 네이밍 규칙 : 소문자 private, 대문자 : public
- Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 x, 프로그램의 시작점 start point

___

## 접근제어 및 Alias

> [access1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/access1.go)
> 
> [access2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/access2.go)

- 패키지 접근제어
- 변수, 상수, 함수, 메서드, 구조체 등 식별자
- 첫글자가 대문자 : 패키지 외부에서 접근가능
- 첫글자가 소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근가능)
- 별칭 사용
- 빈 식별자 사용

___

## 초기화 메소드(init)
> [init1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init1.go)
> 
> [init2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init2.go)
> 
> [init3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section4/init3.go)

- init : 패키지 로드 시에 가장 먼저 호출된다.
- 가장 먼저 초기화 되는 작업 작성 시 유용하다.
- init()는 여러개 있어도 컴파일된다. 
- 다른 패키지에 있으면 해당 패키지에 있는 init함수가 가장 먼지 실행된다.
![image](https://user-images.githubusercontent.com/20445415/204508694-08e23dcc-76d6-41c1-86d0-0cb1147e8b07.png)

___

## 데이터 타입

### Bool

> [bool1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/bool1.go)
> 
> [bool2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/bool2.go)

- Bool(Boolean) : 참과 거짓
- 조건부 논리 연산자랑 주로 사용 : !, ||(or), &&(and)
- 암묵적 형변환이 안된다. : 0, Nil -> False 변환 없음

### 숫자형 기초

> [numeric1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric1.go)
>
> [numeric2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric2.go)
>
> [numeric3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/numeric3.go)

- 데이터 타입 : 숫자형
- 정수, 실수, 복소수
- 32bit, 64bit, unsigned(양수)
- 정수 : 8진수(0), 16진수(0x), 10진수

### 숫자형 연산

> [number_operation1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation1.go)
>
> [number_operation2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation2.go)
>
> [number_operation3.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/number_operation3.go)

- 숫자 연산(산술, 비교)
- 타입이 같아야 가능하다.
- 다른 타입끼리는 반드시 형 변환 후 연산
- 형 변환이 없을 경우 예외(에러) 발생
- +, -, *, %, /, <<, >>, &, ^

### 문자열 기초

> [string1.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string1.go)
>
> [string2.go](https://github.com/FDongFDong/go_language_practice/blob/main/src/section5/string2.go)

- 큰 따옴표로 감싼다 "" 또는 백스쿼트 ``
- Golang : 문자 char 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
- 문자 : ''로 작성 가능
- 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)...
- 문자열 : UTF-8 인코딩(유니코드 문자 집합) -> 바이트 수 주의

### 문자열 연산

> [string_oper1.go]()
>
> [string_oper2.go]()
>
> [string_oper3.go]()

- 문자열 연산
- 추출, 비교, 조합(결합)

___

## 자료구조

### 배열

> [array1.go]()
> [array2.go]()
> [array3.go]()

- 배열은 용량, 길이가 항상 같다.
- 배열 vs 슬라이스 차이점 중요
- 길이 고정 vs 길이 가변
- 값 타입 vs 참조 타입
- 복사 전달 vs 참조 값 전달
- 전체 비교연산자 사용가능 vs 비교 연산자 사용불가
- 대부분 슬라이스 사용한다.
- cap() : 배열, 슬라이스 용량 계산
- len() : 배열, 슬라이스 개수 계산





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
