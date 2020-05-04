package elktools_elasticsearch

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/stretchr/testify/assert"
)

func (s *ESTestSuite) TestCreateSaveDeleteIlmPolicy() {

	// Create ILM policy
	_, err := createILMPolicy("test", "../fixtures/ilm/ilm.json", s.client)
	assert.NoError(s.T(), err)

	// Create all ILM policies
	_, err = createAllILMPolicies("../fixtures/ilm", s.client)
	assert.NoError(s.T(), err)

	// Save ILM policy
	_, err = saveIlmPolicy("test", "/tmp/test.json", s.client)
	assert.NoError(s.T(), err)

	// Save all ILM policies
	_, err = saveAllILMPolicies("/tmp", s.client)
	assert.NoError(s.T(), err)

	// Get the ILM status
	req := &esapi.IndicesCreateRequest{
		Index: "test",
		Body:  strings.NewReader(`{"settings": {"index.lifecycle.name": "test"}}`),
	}
	_, err = req.Do(context.Background(), s.client)
	if err != nil {
		panic(err)
	}
	_, err = getStatusILMPOlicy("test", s.client)
	assert.NoError(s.T(), err)

	// Delete ILM policy
	req2 := &esapi.IndicesDeleteRequest{
		Index: []string{"test"},
	}
	_, err = req2.Do(context.Background(), s.client)
	if err != nil {
		panic(err)
	}
	_, err = deleteILMPolicy("test", s.client)
	assert.NoError(s.T(), err)

}

func (s *ESTestSuite) TestStartStopILMService() {

	err := stopILMService(s.client)
	assert.NoError(s.T(), err)

	err = startILMService(s.client)
	assert.NoError(s.T(), err)

}
