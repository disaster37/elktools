package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestCreateIndice() {

	// Create indice
	_, err := createIndice("test", "../fixtures/indice/indice.json", s.client)
	assert.NoError(s.T(), err)

}
