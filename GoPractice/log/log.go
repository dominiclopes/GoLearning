package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIKey = key(42)

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIKey).(int64)
	if !ok {
		log.Println("could not find request ID in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIKey, id)
		f(rw, req.WithContext(ctx))
	}
}
