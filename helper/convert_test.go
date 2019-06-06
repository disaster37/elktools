package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonToBytes(t *testing.T) {
	res, err := JsonToBytes(`{"name": "test"}`)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, "\"{\\\"name\\\": \\\"test\\\"}\"", string(res))
}

func TestBytesToJson(t *testing.T) {
	var res map[string]string
	err := BytesToJson([]byte(`{"name": "test"}`), &res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, "test", res["name"])
}
