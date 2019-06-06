package helper

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {

	err := WriteFile("/tmp/test", []byte("test"))
	assert.NoError(t, err)

	b, err := ioutil.ReadFile("/tmp/test")
	assert.NoError(t, err)
	assert.Equal(t, "test", string(b))
}

func TestListFilesInPath(t *testing.T) {

	res, err := ListFilesInPath("../fixtures", ".json")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Contains(t, res, "../fixtures/ilm.json")
}
