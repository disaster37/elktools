package elktools_kibana

import (
	"os"
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
	address := os.Getenv("KIBANA_URL")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	if address == "" {
		panic("You need to put kibana url on environment variable KIBANA_URL. If you need auth, you can use ELASTICSEARCH_USERNAME and ELASTICSEARCH_PASSWORD")
	}
	cfg := kibana.Config{
		Address:          address,
		Username:         username,
		Password:         password,
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
