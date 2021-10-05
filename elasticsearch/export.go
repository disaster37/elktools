package elktools_elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// ExportDataToFiles permit to extract some datas to files
// It return error if something wrong
func ExportDataToFiles(c *cli.Context) error {

	es, err := manageElasticsearchGlobalParameters(c)
	if err != nil {
		return err
	}

	from := c.String("from")
	to := c.String("to")
	dateField := c.String("date-field")
	index := c.String("index")
	query := c.String("query")
	fields := c.StringSlice("fields")
	separator := c.String("separator")
	splitFileField := c.String("split-file-field")
	path := c.String("path")

	if path == "" {
		return errors.New("You must set --path")
	}
	if splitFileField == "" {
		return errors.New("You must set --split-file-field")
	}
	if query == "" {
		return errors.New("You must set --query")
	}

	err = exportDataToFiles(from, to, dateField, index, query, fields, separator, splitFileField, path, es)
	if err != nil {
		return err
	}

	log.Infof("Extract successfully")

	return nil
}

func exportDataToFiles(fromDate string, toDate string, dateField string, index string, query string, fields []string, separator string, splitFileColumn string, path string, es *elasticsearch.Client) error {

	if path == "" {
		return errors.New("You must provide path")
	}
	if index == "" {
		return errors.New("You must provide index")
	}

	if dateField == "" {
		return errors.New("You must provide date-field")
	}

	if es == nil {
		return errors.New("You must provide es client")
	}

	sortFields := []string{fmt.Sprintf("%s:asc", dateField)}

	log.Debugf("fromDate: %s", fromDate)
	log.Debugf("toDate: %s", toDate)
	log.Debugf("dateField: %s", dateField)
	log.Debugf("index: %s", index)
	log.Debugf("query: %s", query)
	log.Debugf("fields: %s", fields)
	log.Debugf("sortFields: %s", sortFields)
	log.Debugf("separator: %s", separator)
	log.Debugf("splitFileColumn: %s", splitFileColumn)
	log.Debugf("path: %s", path)

	// Build query
	var rangeDateQuery *elastic.RangeQuery

	rangeDateQuery = elastic.NewRangeQuery(dateField).
		Gte(fromDate).
		Lte(toDate)

	stringQuery := elastic.NewQueryStringQuery(query).
		AnalyzeWildcard(true)
	boolQuery := elastic.NewBoolQuery().Must(rangeDateQuery, stringQuery)
	searchRequest := elastic.NewSearchRequest().
		Query(boolQuery)

	searchQuery, err := searchRequest.Body()
	if err != nil {
		return err
	}

	log.Debugf("Query: \n%s", searchQuery)

	computedFields := append(fields, splitFileColumn)

	res, err := es.API.Search(
		es.API.Search.WithContext(context.Background()),
		es.API.Search.WithPretty(),
		es.API.Search.WithIndex(index),
		es.API.Search.WithSort(sortFields...),
		es.API.Search.WithSourceIncludes(computedFields...),
		es.API.Search.WithSize(10000),
		es.API.Search.WithBody(strings.NewReader(searchQuery)),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when extract data: %s", res.String())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	searchResult := &elastic.SearchResult{}
	if err = json.Unmarshal(body, searchResult); err != nil {
		return err
	}

	log.Debugf("Read %d docs", searchResult.TotalHits())

	// Loop over results
	if searchResult.TotalHits() > 0 {
		var data map[string]interface{}
		listFiles := make(map[string]*os.File, 0)

		for _, item := range searchResult.Hits.Hits {
			log.Debugf("Item %s", item.Source)

			// Create target file to write result
			err := json.Unmarshal(item.Source, &data)
			if err != nil {
				return err
			}

			fileName := fmt.Sprintf("%s/%s", path, data[splitFileColumn])
			file, ok := listFiles[fileName]
			if !ok {
				log.Infof("Create file %s", fileName)
				file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer file.Close()
				if err != nil {
					return err
				}
				listFiles[fileName] = file
			}
			// Extract needed columns
			td := make([]string, 0)
			for _, field := range fields {
				td = append(td, data[field].(string))
			}

			// Write result
			_, err = file.WriteString(fmt.Sprintf("%s\n", strings.Join(td, separator)))
			if err != nil {
				return err
			}
		}

	}

	return nil
}
