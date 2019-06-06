package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestCreateIlmPolicy() {

	err := createILMPolicy("test", "../fixtures/ilm.json", s.client)
	assert.NoError(s.T(), err)
}

func (s *ESTestSuite) TestSaveIlmPolicy() {
	err := saveIlmPolicy("test", "/tmp/test.json", s.client)
	assert.NoError(s.T(), err)
}
