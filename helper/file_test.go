package helper

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {

	err := WriteFile("/tmp/test", []byte("test"))
	assert.NoError(t, err)

	b, err := ioutil.ReadFile("/tmp/test")
	assert.NoError(t, err)
	assert.Equal(t, "test", string(b))
}

func TestListFilesInPath(t *testing.T) {

	res, err := ListFilesInPath("../fixtures/ilm", ".json")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Contains(t, res, "../fixtures/ilm/ilm.json")
}

func TestLoadYaml(t *testing.T) {
	data := make(map[interface{}]interface{})
	err := LoadYaml("../fixtures/helper/test.yaml", &data)
	assert.NoError(t, err)
	assert.Equal(t, "test", data["name"])
}
