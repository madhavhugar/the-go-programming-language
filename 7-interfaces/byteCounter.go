package interfaces

// ByteCounter implements the io.Writer interface and
// tracks the total bytes written to it
type ByteCounter int

func (bc *ByteCounter) Write(text []byte) (int, error) {
	*bc += ByteCounter(len(text))
	return len(text), nil
}
