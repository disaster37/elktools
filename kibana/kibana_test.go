package elktools_kibana

import (
	"testing"
	"time"

	"github.com/disaster37/go-kibana-rest/v7"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type ESTestSuite struct {
	suite.Suite
	client *kibana.Client
}

func (s *ESTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	// Init client
	cfg := kibana.Config{
		Address:          "http://golang-12-kb:5601",
		Username:         "elastic",
		Password:         "changeme",
		DisableVerifySSL: false,
	}

	client, err := kibana.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// Wait es is online
	isOnline := false
	for isOnline == false {
		_, err := client.API.KibanaSpaces.List()
		if err == nil {
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

	err := checkConnexion(s.client)
	assert.NoError(s.T(), err)
}
