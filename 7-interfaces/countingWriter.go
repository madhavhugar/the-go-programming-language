package interfaces

import "io"

// CountingWriterStorage wraps a io.Writer
// interface to keep a track of the total
// number of bytes written
type CountingWriterStorage struct {
	writer io.Writer
	count  *int
}

func (cw CountingWriterStorage) Write(text []byte) (int, error) {
	*cw.count += len(text)
	cw.writer.Write(text)
	return len(text), nil
}

// CountingWriter given an io.Writer, returns
// a new Writer that wraps the original, and
// a pointer to an int variables that at any
// moment contains the number of bytes written
// to the new Writer
func CountingWriter(w io.Writer) (io.Writer, *int) {
	var counter int
	counter = 0
	cw := CountingWriterStorage{w, &counter}
	return cw, &counter
}
