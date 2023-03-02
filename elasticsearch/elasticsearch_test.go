package elktools_elasticsearch

import (
	"os"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
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
	//address := os.Getenv("ELASTICSEARCH_URL")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	cfg := elasticsearch.Config{
		Username: username,
		Password: password,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// Wait es is online
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

func (s *ESTestSuite) TestCheckConnexion() {

	clientInfo, err := checkConnexion(s.client)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), clientInfo)
}

func (s *ESTestSuite) TestCheckCluster() {

	clusterStatus, err := checkClusterStatus(s.client)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "green", clusterStatus)
}

func (s *ESTestSuite) TestClusterRoutingAllocation() {

	err := disableRoutingAllocation(s.client)
	assert.NoError(s.T(), err)

	err = enableRoutingAllocation(s.client)
	assert.NoError(s.T(), err)
}
