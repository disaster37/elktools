package elktools_elasticsearch

import (
	"context"
	"fmt"
	"github.com/disaster37/elktools/helper"
	"github.com/elastic/go-elasticsearch"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"strings"
)

// CreateILMPolicy permit to create or update Lifecycle policy
// It return error if something wrong
func CreateILMPolicy(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
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

	res, err := es.Ilm.PutLifecycle(
		es.Ilm.PutLifecycle.WithContext(context.Background()),
		es.Ilm.PutLifecycle.WithPretty(),
		es.Ilm.PutLifecycle.WithPolicy(id),
		es.Ilm.PutLifecycle.WithBody(strings.NewReader(policyJson)),
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

	es, err := manageElasticsearchGlobalParameters()
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
	res, err := es.Ilm.GetLifecycle(
		es.Ilm.GetLifecycle.WithContext(context.Background()),
		es.Ilm.GetLifecycle.WithPretty(),
		es.Ilm.GetLifecycle.WithPolicy(id),
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

	es, err := manageElasticsearchGlobalParameters()
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

	res, err := es.Ilm.DeleteLifecycle(
		es.Ilm.DeleteLifecycle.WithContext(context.Background()),
		es.Ilm.DeleteLifecycle.WithPretty(),
		es.Ilm.DeleteLifecycle.WithPolicy(id),
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

	es, err := manageElasticsearchGlobalParameters()
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

func saveAllILMPolicies(path string, es *elasticsearch.Client) (map[string]interface{}, error) {
	if path == "" {
		return nil, errors.New("You must set path")
	}

	if es == nil {
		return nil, errors.New("You must set es")
	}

	// Read the policy from API
	res, err := es.Ilm.GetLifecycle(
		es.Ilm.GetLifecycle.WithContext(context.Background()),
		es.Ilm.GetLifecycle.WithPretty(),
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

	es, err := manageElasticsearchGlobalParameters()
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
func GetStatusIlmPolicy(c *cli.Context) error {
	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	elasticsearchIndex := c.String("elasticsearch-index")

	if elasticsearchIndex == "" {
		return errors.New("You must set elasticsearch-index")
	}

	log.Debugf("Elasticsearch index: %s", elasticsearchIndex)

	res, err := es.Ilm.ExplainLifecycle(
		es.Ilm.ExplainLifecycle.WithContext(context.Background()),
		es.Ilm.ExplainLifecycle.WithPretty(),
		es.Ilm.ExplainLifecycle.WithIndex(elasticsearchIndex),
	)

	defer res.Body.Close()

	if res.IsError() {
		errors.Errorf("Error when get lifecycle policy status on index %s: %s", elasticsearchIndex, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Infof("Lifecycle policy status on index %s:\n%s", elasticsearchIndex, body)

	if err != nil {
		return err
	}

	return nil
}
