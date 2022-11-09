package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTimeInterval(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		res := GenerateTimeInterval(5)
		assert.Equal(t, 288, len(res))
	})
}
