// 최종 실습 예제
// 대상 사이트 : 네이버 랭킹 뉴스(https://news.nate.com/rank/?mid=n1000)

package main

import (
	"fmt"
	"net/http"
	"sync"

	_ "github.com/yhat/scrape"
	"golang.org/x/net/html"
	_ "golang.org/x/net/html/atom"
)

const (
	urlRoot = "https://news.nate.com/rank/?mid=n100"
)

// yhat/scrape : 사용하기 어렵지만, 코드 학습 위해서 사용
// goquery : goquery 기반 강력하고 쉬운 패키지(가장 많이 사용)
// PuerkitoBio/goquery : 쉬운 HTML Parsing 지원

// 에러 발생
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

var wg sync.WaitGroup

func main() {
	// 메인 페이지 Get 요청
	response, err := http.Get(urlRoot)
	errCheck(err)

	// 요청 Body 닫기
	defer response.Body.Close()

	// 응답 데이터, Body 내용 파싱
	root, err := html.Parse(response.Body)
	errCheck(err)

	fmt.Println(root)

}
