package elktools_elasticsearch

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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

func CheckConnexion(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	clientInfo, err := checkConnexion(es)
	if err != nil {
		return err
	}

	log.Infof("Connexion OK:\n%s", clientInfo)

	return nil
}

func checkConnexion(es *elasticsearch.Client) (string, error) {

	res, err := es.Info()
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when check Elasticsearch connexion: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
