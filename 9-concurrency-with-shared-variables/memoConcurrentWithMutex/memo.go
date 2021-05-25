package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type result struct {
	resp interface{}
	err error
}

type entry struct {
	result result
	ready chan struct{}
}

func getHTTPBody(url string) (interface{}, error) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}

type Memo struct {
	f func (key string) (interface{}, error)
	cached map[string]*entry
	mu sync.Mutex
}

func New() *Memo {
	return &Memo{
		f:      getHTTPBody,
		cached: make(map[string]*entry),
		mu:     sync.Mutex{},
	}
}

func (m *Memo) Get(url string) (interface{}, error) {
	m.mu.Lock()
	e := m.cached[url]

	if e == nil {
		r, err := m.f(url)
		ent := &entry{
			result: result{resp: r, err:  err},
			ready:  make(chan struct{}),
		}
		e = ent
		m.mu.Unlock()
		close(e.ready)
	} else {
		m.mu.Unlock()
		<- e.ready
	}

	return e.result.resp, e.result.err
}
