package elktools_elasticsearch

import (
	"context"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func EnableMlUpgradeMode(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = enableMlUpgradeMode(es)
	if err != nil {
		return err
	}

	log.Infof("Enable upgrade mode on ML successfully")

	return nil

}

func DisableMlUpgradeMode(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = disableMlUpgradeMode(es)
	if err != nil {
		return err
	}

	log.Infof("Disable upgrade mode on ML successfully")

	return nil

}

func enableMlUpgradeMode(es *elasticsearch.Client) error {
	res, err := es.ML.SetUpgradeMode(
		es.ML.SetUpgradeMode.WithContext(context.Background()),
		es.ML.SetUpgradeMode.WithEnabled(true),
		es.ML.SetUpgradeMode.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when enable ML upgrade mode: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}

func disableMlUpgradeMode(es *elasticsearch.Client) error {
	res, err := es.ML.SetUpgradeMode(
		es.ML.SetUpgradeMode.WithContext(context.Background()),
		es.ML.SetUpgradeMode.WithEnabled(false),
		es.ML.SetUpgradeMode.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when enable ML upgrade mode: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}
