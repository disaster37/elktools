package elktools_elasticsearch

import (
	"context"
	"fmt"
	"github.com/disaster37/elktools/helper"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"strings"
	"github.com/pkg/errors"
	"github.com/elastic/go-elasticsearch"
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
		return cli.NewExitError("You must set lifecycle-policy-id", 1)
	}
	if lifecyclePolicyFile == "" {
		return cli.NewExitError("You must set lifecycle-policy-file parameter", 1)
	}

	err = createILMPolicy(lifecyclePolicyID, lifecyclePolicyFile, es)
	if err != nil {
		return err
	}
	
	return nil
}

// createILMPolicy permit to create or update lifecycle policy on Elasticsearch from file
func createILMPolicy(id string, file string, es *elasticsearch.Client) error {

	if id == "" {
		errors.New("You must provide id")
	}
	if file == "" {
		return errors.New("You must provide file")
	}

	log.Debug("ID: ", id)
	log.Debug("File: ", file)

	// Read the policy file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
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
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
	    errors.Errorf("Error when add lifecycle policy %s: %s", id, res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Infof("Add life cycle policy %s successfully:\n%s", id, body)

	return nil
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
		return cli.NewExitError("You must set lifecycle-policy-id", 1)
	}
	if lifecyclePolicyFile == "" {
		return cli.NewExitError("You must set lifecycle-policy-file parameter", 1)
	}

	log.Debug("Lifecycle policy ID: ", lifecyclePolicyID)
	log.Debug("Lifecycle policy file: ", lifecyclePolicyFile)

	// Read the policy from API
	res, err := es.Ilm.GetLifecycle(
		es.Ilm.GetLifecycle.WithContext(context.Background()),
		es.Ilm.GetLifecycle.WithPretty(),
		es.Ilm.GetLifecycle.WithPolicy(lifecyclePolicyID),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return cli.NewExitError(fmt.Sprintf("Error when get lifecycle policy %s: %s", lifecyclePolicyID, res.String()), 1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Debugf("Get life cycle policy %s successfully:\n%s", lifecyclePolicyID, body)

	// Write contend to file
	err = helper.WriteFile(lifecyclePolicyFile, body)
	if err != nil {
		return err
	}

	log.Infof("Write life cycle policy %s successfully on file %s", lifecyclePolicyID, lifecyclePolicyFile)

	return nil
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
		return cli.NewExitError("You must set lifecycle-policy-id", 1)
	}

	log.Debug("Lifecycle policy ID: ", lifecyclePolicyID)

	res, err := es.Ilm.DeleteLifecycle(
		es.Ilm.DeleteLifecycle.WithContext(context.Background()),
		es.Ilm.DeleteLifecycle.WithPretty(),
		es.Ilm.DeleteLifecycle.WithPolicy(lifecyclePolicyID),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return cli.NewExitError(fmt.Sprintf("Error when delete lifecycle policy %s: %s", lifecyclePolicyID, res.String()), 1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Infof("Delete life cycle policy %s successfully:\n%s", lifecyclePolicyID, body)

	return nil
}

// SaveAllILMPolices permit to save all lifecycle policy
func SaveAllILMPolicies(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	lifecyclePolicyBasePath := c.String("lifecycle-policy-base-path")

	if lifecyclePolicyBasePath == "" {
		return cli.NewExitError("You must set lifecycle-policy-base-path", 1)
	}

	log.Debug("Lifecycle policy base path: ", lifecyclePolicyBasePath)

	// Read the policy from API
	res, err := es.Ilm.GetLifecycle(
		es.Ilm.GetLifecycle.WithContext(context.Background()),
		es.Ilm.GetLifecycle.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return cli.NewExitError(fmt.Sprintf("Error when get all lifecycle policies: %s", res.String()), 1)
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Debugf("Get all life cycle policies successfully:\n%s", body)
	if err != nil {
		return err
	}

	var r map[string]interface{}
	err = helper.BytesToJson(body, &r)
	if err != nil {
		return err
	}

	for name, lifeCyclePolicy := range r {

		log.Debugf("Process the lifecycle policy: %s", name)

		data, err := helper.JsonToBytes(lifeCyclePolicy)
		if err != nil {
			return err
		}
		err = helper.WriteFile(fmt.Sprintf("%s/%s.json", lifecyclePolicyBasePath, name), data)
		if err != nil {
			return err
		}

		log.Infof("Save lifecycle policy %s on %s successfully", name, fmt.Sprintf("%s/%s.json", lifecyclePolicyBasePath, name))
	}

	return nil
}

// CreateAllILMPolices permit to create all lifecycle policies
func CreateAllILMPolicies(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters()
	if err != nil {
		return err
	}

	lifecyclePolicyBasePath := c.String("lifecycle-policy-base-path")

	if lifecyclePolicyBasePath == "" {
		return cli.NewExitError("You must set lifecycle-policy-base-path", 1)
	}

	log.Debug("Lifecycle policy base path: ", lifecyclePolicyBasePath)

	// Read lifecycle policy from file and create it on Elasticsearch
	files, err := helper.ListFilesInPath(lifecyclePolicyBasePath, ".json")
	if err != nil {
		return err
	}

	for _, file := range files {

		// Extract the policy name from the file name
		match, err := helper.ExtractFromRegex("([^\\/]+)\\.json$", file)
		if match == nil {
			return errors.Errorf("Can't extract the lifecycle policy id from the file %s", file)
		}

		err = createILMPolicy(match[1], file, es)
		if err != nil {
			return err
		}

	}

	return nil

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
