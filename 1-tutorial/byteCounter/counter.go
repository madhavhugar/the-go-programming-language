package byteCounter

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func counter(filename string) (int64, error) {
	r, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file %s", err.Error())
		return 0, err
	}

	n, err := io.Copy(ioutil.Discard, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error copying file %s", err.Error())
		return 0, err
	}

	return n, nil
}
