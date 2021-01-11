package interfaces

import (
	"bufio"
	"strings"
)

// WordCounter implements the io.Writer interface
// and counts the total number of words written
type WordCounter int

func (wc *WordCounter) Write(text []byte) {
	s := strings.Split(string(text), " ")
	*wc += WordCounter(len(s))
}

// WordCounterBufio implements io.Write
// and uses bufio methods to count the
// total number of words written
type WordCounterBufio int

func (wcb *WordCounterBufio) Write(text []byte) {
	s := string(text)
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var count int
	for scanner.Scan() {
		count++
	}
	*wcb += WordCounterBufio(count)
}
