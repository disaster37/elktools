package elktools_elasticsearch

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// CreateIndice permit to create new indice with settings
func CreateIndice(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	indiceSettingFile := c.String("indice-setting-file")
	indiceName := c.String("indice-name")
	if indiceName == "" {
		return errors.New("You must set indice-name parameter")
	}
	if indiceSettingFile == "" {
		return errors.New("You must set indice-setting-file parameter")
	}
	log.Debug("Indice name: ", indiceName)
	log.Debug("Indice setting file: ", indiceSettingFile)

	b, err := createIndice(indiceName, indiceSettingFile, es)
	if err != nil {
		return err
	}

	log.Infof("Indice %s created successfully:\n%s", indiceName, b)

	return nil

}

// createIndice permit to create new indice with settings from file
func createIndice(name string, file string, es *elasticsearch.Client) (string, error) {
	if name == "" {
		return "", errors.New("You must set the name")
	}
	if file == "" {
		return "", errors.New("You must set file")
	}
	if es == nil {
		return "", errors.New("You must set es")
	}
	log.Debug("Name: ", name)
	log.Debug("file: ", file)

	// Read the settings file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	configJson := string(b)
	log.Debug("Indice settings: ", configJson)

	res, err := es.API.Indices.Create(
		name,
		es.API.Indices.Create.WithBody(strings.NewReader(configJson)),
		es.API.Indices.Create.WithContext(context.Background()),
		es.API.Indices.Create.WithPretty(),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when create indice %s: %s", name, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
