package elktools_elasticsearch

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"strings"
)

func manageElasticsearchGlobalParameters(c *cli.Context) (*elasticsearch.Client, error) {

	if c.GlobalString("url") == "" {
		return nil, errors.New("You must set --url parameter")
	}

	log.Debug("Elasticsearch URL: ", c.GlobalString("url"))
	log.Debug("Elasticsearch user: ", c.GlobalString("user"))
	log.Debug("Elasticsearch password: XXX")
	log.Debug("Disable verify SSL: ", c.GlobalBool("self-signed-certificate"))

	// Init es client
	elasticsearchURLs := strings.Split(c.GlobalString("url"), ",")
	cfg := elasticsearch.Config{
		Addresses: elasticsearchURLs,
		Username:  c.GlobalString("user"),
		Password:  c.GlobalString("password"),
	}
	if c.GlobalBool("self-signed-certificate") == true {
		cfg.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil

}
