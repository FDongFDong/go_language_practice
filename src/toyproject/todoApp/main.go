package main

import (
	"net/http"
	"todoApp/app"

	"github.com/urfave/negroni"
)

func main() {

	m := app.MakeNewHandler()
	// negroni : 로그와 파일서버를 제공해주는 미들웨어
	n := negroni.Classic()
	n.UseHandler(m)

	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
