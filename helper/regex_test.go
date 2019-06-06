package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractFromRegex(t *testing.T) {

	res, err := ExtractFromRegex("([^\\/]+).json$", "/tmp/test.json")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, "test", res[1])
}
