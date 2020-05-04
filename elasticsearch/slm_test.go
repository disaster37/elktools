package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestStartStopSLMService() {

	err := stopSLMService(s.client)
	assert.NoError(s.T(), err)

	err = startSLMService(s.client)
	assert.NoError(s.T(), err)

}
