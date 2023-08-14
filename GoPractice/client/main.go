package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/home", nil)
	if err != nil {
		log.Fatalf("error creating request, %v", err)
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error executing request, %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("incorrect status code, observed: %v, expected: 200", resp.StatusCode)
	}

	io.Copy(os.Stdout, resp.Body)
}
