# go_language_practice

- [go\_language\_practice](#go_language_practice)
  - [Go 언어 장점 및 특징](#go-언어-장점-및-특징)
  - [데이터 타입(자료형) + 변수를 먼저 배우는 이유](#데이터-타입자료형--변수를-먼저-배우는-이유)
  - [변수 상수](#변수-상수)
  - [포인터](#포인터)
  - [함수](#함수)
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

## 데이터 타입(자료형) + 변수를 먼저 배우는 이유

- 데이터 저장/조회(관리)
- 연산

의 근간이 되기 떄문이다.

## 변수 상수

> [type.go](https://github.com/FDongFDong/go_language_practice/blob/main/type/type.go)
___

## 포인터

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
