package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestStartStopWatcherService() {

	err := stopWatcherService(s.client)
	assert.NoError(s.T(), err)

	err = startWatcherService(s.client)
	assert.NoError(s.T(), err)

}
