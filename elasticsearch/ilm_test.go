package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestCreateSaveDeleteIlmPolicy() {

	// Create ILM policy
	_, err := createILMPolicy("test", "../fixtures/ilm.json", s.client)
	assert.NoError(s.T(), err)

	// Create all ILM policies
	_, err = createAllILMPolicies("../fixtures", s.client)
	assert.NoError(s.T(), err)

	// Save ILM policy
	_, err = saveIlmPolicy("test", "/tmp/test.json", s.client)
	assert.NoError(s.T(), err)

	// Save all ILM policies
	_, err = saveAllILMPolicies("/tmp", s.client)
	assert.NoError(s.T(), err)

	// Delete ILM policy
	_, err = deleteILMPolicy("test", s.client)
	assert.NoError(s.T(), err)

}
