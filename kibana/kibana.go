package elktools_kibana

import (
	"github.com/disaster37/go-kibana-rest"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	KibanaUrl        string
	User             string
	Password         string
	DisableVerifySSL bool
)

// Check the global parameter
func manageKibanaGlobalParameters() (*kibana.Client, error) {

	if KibanaUrl == "" {
		return nil, errors.New("You must set --kibana-url parameter")
	}

	log.Debug("Kibana URL: ", KibanaUrl)
	log.Debug("Kibana user: ", User)
	log.Debug("Kibana password: XXX")
	log.Debug("Disable verify SSL: ", DisableVerifySSL)

	// Init kibana client
	cfg := kibana.Config{
		Address:          KibanaUrl,
		Username:         User,
		Password:         Password,
		DisableVerifySSL: DisableVerifySSL,
	}

	kb, err := kibana.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return kb, nil

}
