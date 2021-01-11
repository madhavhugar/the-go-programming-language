package interfaces

import (
	"bufio"
	"strings"
)

// LineCounterBufio implements the io.Writer interface
// and uses bufio.Scan methods to count the total
// number of lines written
type LineCounterBufio int

func (lc *LineCounterBufio) Write(text []byte) {
	s := string(text)
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var count int
	for scanner.Scan() {
		count++
	}
	*lc += LineCounterBufio(count)
}
