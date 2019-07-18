package elktools_elasticsearch

import (
	"context"
	"github.com/disaster37/elktools/helper"
	"github.com/elastic/go-elasticsearch"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"strings"
)

// CreateIndiceTemplate permit to create or update from cli the indice template
func CreateIndiceTemplate(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	indiceTemplateFile := c.String("indice-template-file")
	indiceTemplateID := c.String("indice-template-id")
	if indiceTemplateID == "" {
		return errors.New("You must set indice-template-id parameter")
	}
	if indiceTemplateFile == "" {
		return errors.New("You must set indice-template-file parameter")
	}
	log.Debug("Indice template ID: ", indiceTemplateID)
	log.Debug("Indice template file: ", indiceTemplateFile)

	b, err := createIndiceTemplate(indiceTemplateID, indiceTemplateFile, es)
	if err != nil {
		return err
	}

	log.Infof("Indice template %s created successfully:\n%s", indiceTemplateID, b)

	return nil

}

// createIndiceTemplae permit to create or update indice template
func createIndiceTemplate(id string, file string, es *elasticsearch.Client) (string, error) {
	if id == "" {
		return "", errors.New("You must set the id")
	}
	if file == "" {
		return "", errors.New("You must set file")
	}
	if es == nil {
		return "", errors.New("You must set es")
	}
	log.Debug("ID: ", id)
	log.Debug("file: ", file)

	// Read the template file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	templateJson := string(b)
	log.Debug("Template: ", templateJson)

	res, err := es.API.Indices.PutTemplate(
		strings.NewReader(templateJson),
		id,
		es.API.Indices.PutTemplate.WithContext(context.Background()),
		es.API.Indices.PutTemplate.WithPretty(),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when add template %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// CreateAllIndiceTemplates permit to create or update from cli all indice templates found on the given folder
func CreateAllIndiceTemplates(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	indiceTemplatePath := c.String("indice-template-path")
	if indiceTemplatePath == "" {
		return errors.New("You must set indice-template-path parameter")
	}
	log.Debug("Indice template path: ", indiceTemplatePath)

	res, err := createAllIndiceTemplates(indiceTemplatePath, es)
	if err != nil {
		return err
	}

	log.Infof("%d Indices templates created successfully from path %s", len(res), indiceTemplatePath)

	return nil

}

// createAllIndiceTemplates permet to create or update all indice template from file found on givent path
func createAllIndiceTemplates(path string, es *elasticsearch.Client) ([]string, error) {
	if path == "" {
		return nil, errors.New("You must set path")
	}
	if es == nil {
		return nil, errors.New("You must set es")
	}
	log.Debug("Path:", path)

	//List file on path and iterate over them
	listFiles, err := helper.ListFilesInPath(path, ".json")
	if err != nil {
		return nil, err
	}
	for _, file := range listFiles {

		// Extract the indice template name from the file name
		match, err := helper.ExtractFromRegex("([^\\/]+)\\.json$", file)
		if match == nil {
			return nil, errors.Errorf("Can't extract the indice template id from the file %s", file)
		}

		body, err := createIndiceTemplate(match[1], file, es)
		if err != nil {
			return nil, err
		}

		log.Infof("Add indice template %s successfully:\n%s", match[1], body)

	}

	return listFiles, nil

}

// DeleteIndiceTemplate permit to delete indice template from cli
func DeleteIndiceTemplate(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	indiceTemplateID := c.String("indice-template-id")
	if indiceTemplateID == "" {
		return errors.New("You must set indice-template-id parameter")
	}
	log.Debug("Indice template ID: ", indiceTemplateID)

	res, err := deleteIndiceTemplate(indiceTemplateID, es)
	if err != nil {
		return err
	}

	log.Infof("Indice template %s successfully deleted:\n%s", indiceTemplateID, res)

	return nil

}

//deleteIndiceTemplate permit to delete indice template
func deleteIndiceTemplate(id string, es *elasticsearch.Client) (string, error) {
	if id == "" {
		return "", errors.New("You must the id")
	}
	if es == nil {
		return "", errors.New("You must set es")
	}
	log.Debug("ID: ", id)

	res, err := es.API.Indices.DeleteTemplate(
		id,
		es.API.Indices.DeleteTemplate.WithContext(context.Background()),
		es.API.Indices.DeleteTemplate.WithPretty(),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when delete template %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// SaveIndiceTemplate permit to save indice template on file from cli
func SaveIndiceTemplate(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	indiceTemplateID := c.String("indice-template-id")
	indiceTemplateFile := c.String("indice-template-file")
	if indiceTemplateID == "" {
		return errors.New("You must set indice-template-id parameter")
	}
	if indiceTemplateFile == "" {
		return errors.New("You must set indice-template-file parameter")
	}
	log.Debug("Indice template ID: ", indiceTemplateID)
	log.Debug("Indice template file: ", indiceTemplateFile)

	_, err = saveIndiceTemplate(indiceTemplateID, indiceTemplateFile, es)
	if err != nil {
		return err
	}

	log.Infof("Save indice template %s successfully on %s", indiceTemplateID, indiceTemplateFile)

	return nil

}

// saveIndiceTemplate permit to save the given indice template on file
func saveIndiceTemplate(id string, file string, es *elasticsearch.Client) (string, error) {

	if id == "" {
		return "", errors.New("You must the id")
	}
	if file == "" {
		return "", errors.New("You must set file")
	}
	if es == nil {
		return "", errors.New("You must set es")
	}
	log.Debug("ID: ", id)
	log.Debug("File: ", file)

	res, err := es.API.Indices.GetTemplate(
		es.API.Indices.GetTemplate.WithName(id),
		es.API.Indices.GetTemplate.WithContext(context.Background()),
		es.API.Indices.GetTemplate.WithPretty(),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when delete template %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = helper.WriteFile(file, body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
