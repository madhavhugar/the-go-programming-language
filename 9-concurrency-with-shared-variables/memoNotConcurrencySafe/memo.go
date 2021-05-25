package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHTTPBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type result struct {
	resp interface{}
	err error
}

type Memo struct {
	f func(key string) (interface{}, error)
	// cached is being used without a mutex and does not have a monitor goroutine either
	cached map[string]result
}

func New() *Memo {
	return &Memo{
		f:      getHTTPBody,
		cached: make(map[string]result),
	}
}

func (m *Memo) Get(url string) (interface{}, error) {
	cachedResponse, ok := m.cached[url]
	if !ok {
		r, err := getHTTPBody(url)
		m.cached[url] = result{resp: r, err: err}
		return cachedResponse.resp, cachedResponse.err
	}

	return cachedResponse.resp, cachedResponse.err
}
