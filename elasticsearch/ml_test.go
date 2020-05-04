package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestEnableDisableUpgradeML() {

	err := enableMlUpgradeMode(s.client)
	assert.NoError(s.T(), err)

	err = disableMlUpgradeMode(s.client)
	assert.NoError(s.T(), err)

}
