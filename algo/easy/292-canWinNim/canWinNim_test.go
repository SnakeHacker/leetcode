package canWinNim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if i%4 == 0 {
			assert.Equal(t, false, canWinNim(i))
		} else {
			assert.Equal(t, true, canWinNim(i))
		}
	}
}
