package elktools_kibana

import (
	"github.com/disaster37/go-kibana-rest/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Check the global parameter
func manageKibanaGlobalParameters(c *cli.Context) (*kibana.Client, error) {

	if c.GlobalString("url") == "" {
		return nil, errors.New("You must set --url parameter")
	}

	log.Debug("Kibana URL: ", c.GlobalString("url"))
	log.Debug("Kibana user: ", c.GlobalString("user"))
	log.Debug("Kibana password: XXX")
	log.Debug("Disable verify SSL: ", c.GlobalString("self-signed-certificate"))

	// Init kibana client
	cfg := kibana.Config{
		Address:          c.GlobalString("url"),
		Username:         c.GlobalString("user"),
		Password:         c.GlobalString("password"),
		DisableVerifySSL: c.GlobalBool("self-signed-certificate"),
	}

	kb, err := kibana.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return kb, nil

}

func CheckConnexion(c *cli.Context) error {

	kb, err := manageKibanaGlobalParameters(c)
	if err != nil {
		return err
	}

	err = checkConnexion(kb)
	if err != nil {
		return err
	}

	log.Infof("Connexion OK")

	return nil
}

func checkConnexion(kb *kibana.Client) error {
	_, err := kb.API.KibanaSpaces.List()

	return err
}
