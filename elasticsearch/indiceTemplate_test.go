package elktools_elasticsearch

import (
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestCreateIndiceTemplate() {

	// Create indice template
	_, err := createIndiceTemplate("template", "../fixtures/template/template.json", s.client)
	assert.NoError(s.T(), err)

	// Create all indice templates
	_, err = createAllIndiceTemplates("../fixtures/template", s.client)
	assert.NoError(s.T(), err)

	// Save indice template
	_, err = saveIndiceTemplate("template", "/tmp/test.json", s.client)
	assert.NoError(s.T(), err)

	// Delete indice template
	_, err = deleteIndiceTemplate("test", s.client)
	assert.NoError(s.T(), err)

}
