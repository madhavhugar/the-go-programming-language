package taglessSwitch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitchMe(t *testing.T) {
	t.Run("to demonstrate fallback cases do not execute", func(t *testing.T) {
		switchMe(true, true)

		assert.Equal(t, fallbackCaseOne, true)
		assert.Equal(t, fallbackCaseTwo, false)
	})
}
