package elktools_elasticsearch

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/elastic/go-elasticsearch"
	"strings"

)

var (
    ElasticsearchUrl string
    User string
    Password string
)

// Check the global parameter
func manageElasticsearchGlobalParameters() (*elasticsearch.Client, error) {

	if ElasticsearchUrl == "" {
		return nil, errors.New("You must set --elasticsearch-url parameter")
	}
	
    
    log.Debug("Elasticsearch URL: ", ElasticsearchUrl)
    log.Debug("Elasticsearch user: ", User)
    log.Debug("Elasticsearch password: XXX")
    
    
   // Init es client
   elasticsearchURLs := strings.Split(ElasticsearchUrl, ",")
   cfg := elasticsearch.Config{
       Addresses: elasticsearchURLs,
       Username: User,
       Password: Password,
   }
   es, err := elasticsearch.NewClient(cfg)
   if err != nil {
       return nil, err
   }
   
   return es, nil

}