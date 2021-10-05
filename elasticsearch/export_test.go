package elktools_elasticsearch

import (
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestExportDataToFiles() {

	// Exports data without errors
	err := exportDataToFiles("now-1000y", "now", "timestamp", "logs", "*", []string{"message"}, "|", "node.name", "/tmp", s.client)
	assert.NoError(s.T(), err)

	// Check output file exists
	content, err := ioutil.ReadFile("/tmp/es-0")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), content, "[gc][17868238] overhead, spent [334ms] collecting in the last [1s]\n")

	content, err = ioutil.ReadFile("/tmp/es-1")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), content, "[gc][17868264] overhead, spent [279ms] collecting in the last [1s]\n")

}
