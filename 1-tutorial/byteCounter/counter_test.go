package byteCounter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteCounter(t *testing.T) {
	t.Run("byte counter", func(t *testing.T) {
		n, err := counter("./dummy.txt")

		var expected int64 = 12
		assert.Equal(t, expected, n)
		assert.NoError(t, err)
	})
}