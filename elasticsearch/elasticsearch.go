package elktools_elasticsearch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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

	_, err = checkConnexion(es)
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

func CheckClusterStatus(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	status, err := checkClusterStatus(es)
	if err != nil {
		return err
	}

	switch status {
	case "green":
		log.Info("Cluster OK")
		return nil
	case "yellow":
		log.Info("Cluster warning")
		os.Exit(1)
	case "red":
		log.Info("Cluster critical")
		os.Exit(2)
	}

	return nil
}

func checkClusterStatus(es *elasticsearch.Client) (string, error) {
	res, err := es.Cluster.Health(
		es.Cluster.Health.WithPretty(),
	)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when check Elasticsearch status: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var esResponse map[string]interface{}
	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return "", err
	}

	return esResponse["status"].(string), nil

}

func ClusterEnableRoutingAllocation(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = enableRoutingAllocation(es)
	if err != nil {
		return err
	}

	log.Info("Enable routing allocation successfully")

	return nil
}

func ClusterDisableRoutingAllocation(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = disableRoutingAllocation(es)
	if err != nil {
		return err
	}

	log.Info("Disable routing allocation successfully")

	return nil
}

func enableRoutingAllocation(es *elasticsearch.Client) error {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.allocation.enable": "all",
		},
	}

	err := putClusterSettings(es, settings)

	return err
}

func disableRoutingAllocation(es *elasticsearch.Client) error {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.allocation.enable": "primaries",
		},
	}

	err := putClusterSettings(es, settings)

	return err
}

func putClusterSettings(es *elasticsearch.Client, settings map[string]interface{}) error {

	log.Debugf("Settings: %+v", settings)

	data, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	res, err := es.Cluster.PutSettings(
		bytes.NewReader(data),
		es.Cluster.PutSettings.WithPretty(),
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when set Elasticsearch cluster setting: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil

}
