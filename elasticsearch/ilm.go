package elktools_elasticsearch

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/disaster37/elktools/v7/helper"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// CreateILMPolicy permit to create or update Lifecycle policy
// It return error if something wrong
func CreateILMPolicy(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	lifecyclePolicyFile := c.String("lifecycle-policy-file")
	lifecyclePolicyID := c.String("lifecycle-policy-id")

	if lifecyclePolicyID == "" {
		return errors.New("You must set lifecycle-policy-id")
	}
	if lifecyclePolicyFile == "" {
		return errors.New("You must set lifecycle-policy-file parameter")
	}

	body, err := createILMPolicy(lifecyclePolicyID, lifecyclePolicyFile, es)
	if err != nil {
		return err
	}

	log.Infof("Add life cycle policy %s successfully:\n%s", lifecyclePolicyID, body)

	return nil
}

// createILMPolicy permit to create or update lifecycle policy on Elasticsearch from file
func createILMPolicy(id string, file string, es *elasticsearch.Client) (string, error) {

	if id == "" {
		return "", errors.New("You must provide id")
	}
	if file == "" {
		return "", errors.New("You must provide file")
	}

	if es == nil {
		return "", errors.New("You must provide es client")
	}

	log.Debug("ID: ", id)
	log.Debug("File: ", file)

	// Read the policy file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	policyJson := string(b)
	log.Debug("Policy: ", policyJson)

	res, err := es.API.ILM.PutLifecycle(
		es.API.ILM.PutLifecycle.WithContext(context.Background()),
		es.API.ILM.PutLifecycle.WithPretty(),
		es.API.ILM.PutLifecycle.WithPolicy(id),
		es.API.ILM.PutLifecycle.WithBody(strings.NewReader(policyJson)),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when add lifecycle policy %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SaveILMPolicy permit to get and write existing lifecycle policy on file
func SaveILMPolicy(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	lifecyclePolicyFile := c.String("lifecycle-policy-file")
	lifecyclePolicyID := c.String("lifecycle-policy-id")

	if lifecyclePolicyID == "" {
		return errors.New("You must set lifecycle-policy-id")
	}
	if lifecyclePolicyFile == "" {
		return errors.New("You must set lifecycle-policy-file parameter")
	}

	log.Debug("Lifecycle policy ID: ", lifecyclePolicyID)
	log.Debug("Lifecycle policy file: ", lifecyclePolicyFile)

	_, err = saveIlmPolicy(lifecyclePolicyID, lifecyclePolicyFile, es)
	if err != nil {
		return err
	}

	log.Infof("Write life cycle policy %s successfully on file %s", lifecyclePolicyID, lifecyclePolicyFile)

	return nil
}

// saveIlmPolicy permit to get lifecycle policy from elasticsearch and save it on file
func saveIlmPolicy(id string, file string, es *elasticsearch.Client) (string, error) {

	if id == "" {
		return "", errors.New("You must provide id")
	}
	if file == "" {
		return "", errors.New("You must provide file")
	}

	if es == nil {
		return "", errors.New("You must provide es client")
	}

	log.Debug("ID: ", id)
	log.Debug("File: ", file)

	// Read the policy from API
	res, err := es.API.ILM.GetLifecycle(
		es.API.ILM.GetLifecycle.WithContext(context.Background()),
		es.API.ILM.GetLifecycle.WithPretty(),
		es.API.ILM.GetLifecycle.WithPolicy(id),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.IsError() {
		return "", errors.Errorf("Error when get lifecycle policy %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	log.Debugf("Get life cycle policy %s successfully:\n%s", id, body)

	// Write contend to file
	err = helper.WriteFile(file, body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// DeleteILMPolicy permit to delete Lifecycle policy
// It return error if something wrong
func DeleteILMPolicy(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	lifecyclePolicyID := c.String("lifecycle-policy-id")

	if lifecyclePolicyID == "" {
		return errors.New("You must set lifecycle-policy-id")
	}

	log.Debug("Lifecycle policy ID: ", lifecyclePolicyID)

	body, err := deleteILMPolicy(lifecyclePolicyID, es)
	if err != nil {
		return err
	}

	log.Infof("Delete life cycle policy %s successfully:\n%s", lifecyclePolicyID, body)

	return nil
}

// deleteILMPolicy permit to delete the given lifecycle policy in Elasticsearch
func deleteILMPolicy(id string, es *elasticsearch.Client) (string, error) {

	if id == "" {
		return "", errors.New("You must provide id")
	}

	if es == nil {
		return "", errors.New("You must provide es client")
	}

	log.Debug("ID: ", id)

	res, err := es.API.ILM.DeleteLifecycle(
		es.API.ILM.DeleteLifecycle.WithContext(context.Background()),
		es.API.ILM.DeleteLifecycle.WithPretty(),
		es.API.ILM.DeleteLifecycle.WithPolicy(id),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when delete lifecycle policy %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// SaveAllILMPolices permit to save all lifecycle policy
func SaveAllILMPolicies(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	lifecyclePolicyBasePath := c.String("lifecycle-policy-base-path")

	if lifecyclePolicyBasePath == "" {
		return errors.New("You must set lifecycle-policy-base-path")
	}

	log.Debug("Lifecycle policy base path: ", lifecyclePolicyBasePath)

	r, err := saveAllILMPolicies(lifecyclePolicyBasePath, es)
	if err != nil {
		return err
	}

	log.Infof("Save %d lifecycle policies on %s", len(r), lifecyclePolicyBasePath)

	return nil
}

// saveAllILMPolicies permit to save on files all ILM policies
// It return le list of ILM policies as map of strings
func saveAllILMPolicies(path string, es *elasticsearch.Client) (map[string]interface{}, error) {
	if path == "" {
		return nil, errors.New("You must set path")
	}

	if es == nil {
		return nil, errors.New("You must set es")
	}

	// Read the policy from API
	res, err := es.API.ILM.GetLifecycle(
		es.API.ILM.GetLifecycle.WithContext(context.Background()),
		es.API.ILM.GetLifecycle.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, errors.Errorf("Error when get all lifecycle policies: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Debugf("Get all life cycle policies successfully:\n%s", body)
	if err != nil {
		return nil, err
	}

	var r map[string]interface{}
	err = helper.BytesToJson(body, &r)
	if err != nil {
		return nil, err
	}

	for name, lifeCyclePolicy := range r {

		log.Debugf("Process the lifecycle policy: %s", name)

		data, err := helper.JsonToBytes(lifeCyclePolicy)
		if err != nil {
			return nil, err
		}
		err = helper.WriteFile(fmt.Sprintf("%s/%s.json", path, name), data)
		if err != nil {
			return nil, err
		}

		log.Infof("Save lifecycle policy %s on %s/%s.json successfully", name, path, name)
	}

	return r, nil

}

// CreateAllILMPolices permit to create all lifecycle policies
func CreateAllILMPolicies(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	lifecyclePolicyBasePath := c.String("lifecycle-policy-base-path")

	if lifecyclePolicyBasePath == "" {
		return errors.New("You must set lifecycle-policy-base-path")
	}

	log.Debug("Lifecycle policy base path: ", lifecyclePolicyBasePath)

	listFiles, err := createAllILMPolicies(lifecyclePolicyBasePath, es)
	if err != nil {
		return err
	}

	log.Infof("Create %d policies from %s", len(listFiles), lifecyclePolicyBasePath)

	return nil

}

// createAllILMPolicies permit to create all ilm policies found on the given path
// It return the list of policies files that it found
func createAllILMPolicies(path string, es *elasticsearch.Client) ([]string, error) {
	if path == "" {
		return nil, errors.New("You must set path")
	}

	if es == nil {
		return nil, errors.New("You must set es")
	}

	// Read lifecycle policy from file and create it on Elasticsearch
	files, err := helper.ListFilesInPath(path, ".json")
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		// Extract the policy name from the file name
		match, err := helper.ExtractFromRegex("([^\\/]+)\\.json$", file)
		if match == nil {
			return nil, errors.Errorf("Can't extract the lifecycle policy id from the file %s", file)
		}

		body, err := createILMPolicy(match[1], file, es)
		if err != nil {
			return nil, err
		}

		log.Infof("Add life cycle policy %s successfully:\n%s", match[1], body)

	}

	return files, nil

}

// GetStatusIlmPolicy permit to get the current status of lifecycle policy on given index
func GetStatusILMPolicy(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	elasticsearchIndex := c.String("elasticsearch-index")

	if elasticsearchIndex == "" {
		return errors.New("You must set elasticsearch-index")
	}

	log.Debugf("Elasticsearch index: %s", elasticsearchIndex)

	body, err := getStatusILMPOlicy(elasticsearchIndex, es)
	if err != nil {
		return err
	}

	log.Infof("Lifecycle policy status on index %s:\n%s", elasticsearchIndex, body)
	return nil
}

// getSatusILMPolicy permit to explain the ILM policies apply on the given index
func getStatusILMPOlicy(index string, es *elasticsearch.Client) (string, error) {
	if index == "" {
		return "", errors.New("You must set index")
	}

	if es == nil {
		return "", errors.New("You must set es")
	}

	res, err := es.API.ILM.ExplainLifecycle(
		es.API.ILM.ExplainLifecycle.WithContext(context.Background()),
		es.API.ILM.ExplainLifecycle.WithPretty(),
		es.API.ILM.ExplainLifecycle.WithIndex(index),
	)

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.Errorf("Error when get lifecycle policy status on index %s: %s", index, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
