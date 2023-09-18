package elktools_elasticsearch

import "github.com/stretchr/testify/assert"

func (s *ESTestSuite) TestCheckExpectedNodes() {

	// When true
	isExpectedNode, err := checkExpectedNumberNodes(s.client, 1)
	assert.NoError(s.T(), err)
	assert.True(s.T(), isExpectedNode)

	// When false
	isExpectedNode, err = checkExpectedNumberNodes(s.client, 10)
	assert.NoError(s.T(), err)
	assert.False(s.T(), isExpectedNode)
}

func (s *ESTestSuite) TestCheckNodeOnlie() {

	// When true
	isOnline, err := checkNodeOnline(s.client, "elasticsearch", []string{"node_name"})
	assert.NoError(s.T(), err)
	assert.True(s.T(), isOnline)

	// When false
	isOnline, err = checkNodeOnline(s.client, "fake", []string{"node_name"})
	assert.NoError(s.T(), err)
	assert.False(s.T(), isOnline)
}
