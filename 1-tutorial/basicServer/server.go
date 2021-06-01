package basicServer

import (
	"fmt"
	"net/http"
	"sync"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}

var (
	mu sync.Mutex
	count int
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count += 1
	fmt.Fprintf(w, "%d", count)
	mu.Unlock()
}