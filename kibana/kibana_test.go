package elktools_kibana

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"testing"
	"time"
)

type ESTestSuite struct {
	suite.Suite
	client *elasticsearch.Client
}

func (s *ESTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	// Init client
	ElasticsearchUrl = "http://es:9200"
	User = "elastic"
	Password = "changeme"

	// Wait es is online
	client, err := manageElasticsearchGlobalParameters()
	if err != nil {
		panic(err)
	}
	isOnline := false
	for isOnline == false {
		res, err := client.Info()
		if err == nil && res.IsError() == false {
			isOnline = true
		} else {
			time.Sleep(5 * time.Second)
		}
	}

	s.client = client

}

func (s *ESTestSuite) SetupTest() {

	// Do somethink before each test

}

func TestESTestSuite(t *testing.T) {
	suite.Run(t, new(ESTestSuite))
}

func (s *ESTestSuite) TestManageElasticsearchGlobalParameters() {

	client, err := manageElasticsearchGlobalParameters()
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), client)
}
