package interfaces

import "io"

type limiter struct {
	r        io.Reader
	n, limit uint
}

// TODO: implement limit reader
func (lr *limiter) Read(p []byte) (n int, err error) {
	num, err := lr.r.Read(p)
	lr.limit -= uint(num)
	if lr.limit <= 0 {
		return 0, io.EOF
	}
	return num, nil
}

func LimitReader(r io.Reader, n uint) io.Reader {
	return &limiter{r, 0, n}
}
