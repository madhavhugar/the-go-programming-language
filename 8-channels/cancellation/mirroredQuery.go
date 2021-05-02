package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func mirroredQuery() chan string {
	response := make(chan string, 3)
	ctx, cancel := context.WithCancel(context.Background())
	go respondOrClose(ctx, cancel, response)
	go request(ctx, "https://google.com", response)
	go request(ctx, "https://google.de", response)
	go request(ctx, "https://google.in", response)
	return response
}

func respondOrClose(ctx context.Context, cancel context.CancelFunc, response chan string) {
	for {
		select {
		case body := <- response:
			s := fmt.Sprintf("body: %s", body)
			fmt.Println(s)
			cancel()
		case <- ctx.Done():
			fmt.Println("Response received elsewhere...exiting")
			return
		}
	}
}

func request(ctx context.Context, url string, response chan string) {
	var body string

	req, err := http.NewRequestWithContext(ctx,"GET", url, strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if res == nil {
		return
	}

	defer res.Body.Close()

	var p []byte
	_, err = res.Body.Read(p)
	if err != nil {
		fmt.Println(err)
	}
	response <- string(p)
}

func main() {
	mirroredQuery()

	for {}
}
