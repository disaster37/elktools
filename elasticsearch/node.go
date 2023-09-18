package elktools_elasticsearch

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	olivere "github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func CheckNodeOnline(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	isOnline, err := checkNodeOnline(es, c.String("node-name"), c.StringSlice("labels"))
	if err != nil {
		return err
	}

	if isOnline {
		log.Infof("Node %s is on cluster", c.String("node-name"))
		return nil
	}

	log.Warnf("Node %s not yet on cluster", c.String("node-name"))
	os.Exit(1)

	return nil

}

func CheckExpectedNumberNodes(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	isExpected, err := checkExpectedNumberNodes(es, c.Int("number-nodes"))
	if err != nil {
		return err
	}

	if isExpected {
		log.Infof("All nodes in cluster (%d)", c.Int("number-nodes"))
		return nil
	}

	log.Warnf("The are some nodes lost. We expect %d nodes", c.Int("number-nodes"))
	os.Exit(1)

	return nil

}

func checkNodeOnline(es *elasticsearch.Client, nodeName string, labels []string) (bool, error) {
	res, err := es.Nodes.Info(
		es.Nodes.Info.WithContext(context.Background()),
		es.Nodes.Info.WithPretty(),
	)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return false, errors.Errorf("Error when get nodes info: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	nodesInfo := &olivere.NodesInfoResponse{}
	err = json.Unmarshal(body, nodesInfo)
	if err != nil {
		return false, err
	}

	for _, node := range nodesInfo.Nodes {
		if node.Name == nodeName {
			return true, nil
		}

		for _, label := range labels {
			if node.Attributes[label] == nodeName {
				return true, nil
			}
		}
	}

	return false, nil

}

func checkExpectedNumberNodes(es *elasticsearch.Client, nodesNumber int) (bool, error) {
	res, err := es.Nodes.Info(
		es.Nodes.Info.WithContext(context.Background()),
		es.Nodes.Info.WithPretty(),
	)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return false, errors.Errorf("Error when get nodes info: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	nodesInfo := &olivere.NodesInfoResponse{}
	err = json.Unmarshal(body, nodesInfo)
	if err != nil {
		return false, err
	}

	log.Debugf("Found %d nodes in cluster", len(nodesInfo.Nodes))

	return len(nodesInfo.Nodes) == nodesNumber, nil
}
