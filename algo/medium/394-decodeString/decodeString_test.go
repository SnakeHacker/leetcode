package decodeString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	assert.Equal(t, "ssss", decodeString("2[ss]"))
}

func TestDecodeString2(t *testing.T) {
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
}

func TestDecodeString3(t *testing.T) {
	assert.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
}
func TestDecodeString4(t *testing.T) {
	assert.Equal(t, "abccdcdcdxyz", decodeString("abc3[cd]xyz"))
}

func TestDecodeString5(t *testing.T) {
	assert.Equal(t, "leetcode", decodeString("leetcode"))
}
