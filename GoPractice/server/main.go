package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"example.com/log"
)

func main() {
	http.HandleFunc("/home", log.Decorate(HomeHandlerFunc))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func HomeHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, int(42), int64(100))

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(rw, "Welcome to Home Page\n")
	case <-req.Context().Done():
		log.Println(ctx, ctx.Err().Error())
	}
}
