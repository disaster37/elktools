package elktools_elasticsearch

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func StopWatcherService(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = stopWatcherService(es)
	if err != nil {
		return err
	}

	log.Infof("Stop watcher service successfully")

	return nil

}

func StartWatcherService(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	err = startWatcherService(es)
	if err != nil {
		return err
	}

	log.Infof("Start watcher service successfully")

	return nil

}

func stopWatcherService(es *elasticsearch.Client) error {
	res, err := es.Watcher.Stop(
		es.Watcher.Stop.WithContext(context.Background()),
		es.Watcher.Stop.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when stop watcher service: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}

func startWatcherService(es *elasticsearch.Client) error {
	res, err := es.Watcher.Start(
		es.Watcher.Start.WithContext(context.Background()),
		es.Watcher.Start.WithPretty(),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when start watcher service: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Debugf("Response: %s", string(body))

	return nil
}
