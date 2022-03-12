package judgeCircle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeCircle(t *testing.T) {
	actual := judgeCircle("UD")
	assert.True(t, actual)
}

func TestJudgeCircle2(t *testing.T) {
	actual := judgeCircle("LL")
	assert.False(t, actual)
}
