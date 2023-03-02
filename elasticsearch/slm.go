package elktools_elasticsearch

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func StopSLMService(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = stopSLMService(es)
	if err != nil {
		return err
	}

	log.Infof("Stop SLM service successfully")

	return nil

}

func StartSLMService(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = startSLMService(es)
	if err != nil {
		return err
	}

	log.Infof("Start SLM service successfully")

	return nil

}

func stopSLMService(es *elasticsearch.Client) error {

	res, err := es.SlmStop(
		es.SlmStop.WithContext(context.Background()),
		es.SlmStop.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when stop SLM service: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}

func startSLMService(es *elasticsearch.Client) error {
	res, err := es.SlmStart(
		es.SlmStart.WithContext(context.Background()),
		es.SlmStart.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when start SLM service: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}
