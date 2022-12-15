package main

import (
	"CRUD_gin/sub"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	mapi := &http.Server{
		Addr:           ":8080",
		Handler:        sub.Index(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return mapi.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
