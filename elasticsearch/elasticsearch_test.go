package elktools_elasticsearch

import (
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
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
	cfg := elasticsearch.Config{
		Addresses: []string{"http://golang-12-es:9200"},
		Username:  "elastic",
		Password:  "changeme",
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
