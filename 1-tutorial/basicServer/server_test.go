package basicServer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("basicHandler", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", basicHandler)
		s := httptest.NewServer(mux)

		res, err := http.Get(fmt.Sprintf("%s/hello", s.URL))
		if err != nil {
			t.Error(err)
		}

		b , err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		t.Cleanup(func() {
			res.Body.Close()
		})

		assert.Equal(t, string(b), "hello, world")
	})

	t.Run("countHandler", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/count", countHandler)
		s := httptest.NewServer(mux)

		var finalCount int

		makeRequest := func() {
			res, err := http.Get(fmt.Sprintf("%s/count", s.URL))
			if err != nil {
				t.Error(err)
			}
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

			c, err := strconv.Atoi(string(b))
			if err != nil {
				t.Error(err)
			}

			finalCount = c
		}
		makeRequest()
		makeRequest()

		assert.Equal(t, 2, finalCount)
	})
}
