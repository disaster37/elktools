package elktools_elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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

	ctx := context.Background()
	sortFields := []string{fmt.Sprintf("%s:asc", dateField)}
	scrollDuration := 5 * time.Minute
	size := 10000

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

	// Open PIT to scroll over results
	// Working only for ES >= 10
	/*
		res, err := es.API.OpenPointInTime(
			[]string{index},
			es.API.OpenPointInTime.WithKeepAlive("5m"),
			es.API.OpenPointInTime.WithPretty(),
			es.API.OpenPointInTime.WithContext(ctx),
		)
		if err != nil {
			return err
		}
		if res.IsError() {
			return errors.Errorf("Error when extract data: %s", res.String())
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		pitResponse := &elastic.OpenPointInTimeResponse{}
		if err = json.Unmarshal(body, pitResponse); err != nil {
			return err
		}
	*/

	// Build query
	rangeDateQuery := elastic.NewRangeQuery(dateField).
		Gte(fromDate).
		Lte(toDate)
	stringQuery := elastic.NewQueryStringQuery(query).
		AnalyzeWildcard(true)
	boolQuery := elastic.NewBoolQuery().Must(rangeDateQuery, stringQuery)
	searchRequest := elastic.NewSearchRequest().
		Query(boolQuery)
		/*
			PointInTime(&elastic.PointInTime{
				Id:        pitResponse.Id,
				KeepAlive: "5m",
			}).
			SearchAfter()
		*/
	searchQuery, err := searchRequest.Body()
	if err != nil {
		return err
	}
	log.Debugf("Query: \n%s", searchQuery)

	// Forge payload
	computedFields := append(fields, splitFileColumn)
	res, err := es.API.Search(
		es.API.Search.WithContext(ctx),
		es.API.Search.WithPretty(),
		es.API.Search.WithIndex(index),
		es.API.Search.WithSort(sortFields...),
		es.API.Search.WithSourceIncludes(computedFields...),
		es.API.Search.WithSize(size),
		es.API.Search.WithBody(strings.NewReader(searchQuery)),
		es.API.Search.WithScroll(scrollDuration),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("Error when extract data: %s", res.String())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	searchResult := &elastic.SearchResult{}
	if err = json.Unmarshal(body, searchResult); err != nil {
		return err
	}

	log.Infof("Found %d document to export", searchResult.TotalHits())

	isMoreResultToProcess := true
	for isMoreResultToProcess {
		if err = processExport(searchResult, fields, separator, path, splitFileColumn); err != nil {
			return err
		}
		if len(searchResult.Hits.Hits) < size {
			isMoreResultToProcess = false
			log.Debugf("End of scroll. NB doc %d", len(searchResult.Hits.Hits))

			// We clean current scroll
			res, err = es.API.ClearScroll(
				es.API.ClearScroll.WithScrollID(searchResult.ScrollId),
				es.API.ClearScroll.WithContext(ctx),
				es.API.ClearScroll.WithPretty(),
			)
			if err != nil {
				return err
			}
			if res.IsError() {
				return errors.Errorf("Error when clear scroll: %s", res.String())
			}
		} else {
			log.Debugf("Continue with the next scroll")
			res, err := es.API.Scroll(
				es.API.Scroll.WithScrollID(searchResult.ScrollId),
				es.API.Scroll.WithScroll(scrollDuration),
				es.API.Scroll.WithContext(ctx),
				es.API.Scroll.WithPretty(),
			)

			if err != nil {
				return err
			}

			defer res.Body.Close()

			if res.IsError() {
				return errors.Errorf("Error when extract data: %s", res.String())
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			searchResult = &elastic.SearchResult{}
			if err = json.Unmarshal(body, searchResult); err != nil {
				return err
			}
		}
	}

	return nil
}

func processExport(searchResult *elastic.SearchResult, fields []string, separator string, path string, splitFileColumn string) (err error) {

	log.Debugf("Process %d documents", len(searchResult.Hits.Hits))

	// Loop over results
	if len(searchResult.Hits.Hits) > 0 {
		listFiles := make(map[string]*os.File, 0)

		for _, item := range searchResult.Hits.Hits {

			// Create target file to write result
			jsonResult := gjson.ParseBytes(item.Source)

			fileName := fmt.Sprintf("%s/%s", path, jsonResult.Get(splitFileColumn))
			file, ok := listFiles[fileName]
			if !ok {
				if _, err = os.Stat(fileName); os.IsNotExist(err) {
					log.Infof("Create file: %s", fileName)
				}
				log.Debugf("Open file %s", fileName)
				file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Errorf("Error when open file: %s", err.Error())
					return err
				}
				defer file.Close()

				listFiles[fileName] = file
			}
			// Extract needed columns
			td := make([]string, 0)
			for _, field := range fields {
				td = append(td, jsonResult.Get(field).Str)
			}

			// Write result
			_, err := file.WriteString(fmt.Sprintf("%s\n", strings.Join(td, separator)))
			if err != nil {
				log.Errorf("Error when write file: %s", err.Error())
				return err
			}
		}

	}

	return nil
}
