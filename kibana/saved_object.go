package elktools_kibana

import (
	"io/ioutil"

	"github.com/disaster37/elktools/v7/helper"
	"github.com/disaster37/go-kibana-rest/v7"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func exportDashboards(filePath string, userSpace string, kb *kibana.Client) error {
	if filePath == "" {
		return errors.New("You must provide filePath")
	}
	if kb == nil {
		return errors.New("You must provide kb client")
	}
	log.Debug("FilePath: ", filePath)
	log.Debug("UserSpace: ", userSpace)

	// Exports all dashboard and includes all references
	data, err := kb.API.KibanaSavedObject.Export(
		[]string{"dashboard"},
		nil,
		true,
		userSpace,
	)
	if err != nil {
		return err
	}

	// Write all in file
	err = helper.WriteFile(filePath, data)
	if err != nil {
		return err
	}

	return nil
}

// ExportDashboards permit to export all dashboard  and include all references
func ExportDashboards(c *cli.Context) error {
	kb, err := manageKibanaGlobalParameters(c)
	if err != nil {
		return err
	}

	exportFilePath := c.String("file-path")
	userSpace := c.String("user-space")

	if exportFilePath == "" {
		return errors.New("You must set --file-path")
	}
	if userSpace == "" {
		return errors.New("You must set --user-space")
	}

	err = exportDashboards(exportFilePath, userSpace, kb)
	if err != nil {
		return err
	}

	log.Infof("Export all dashboards successfully in %s", exportFilePath)

	return nil
}

func importDashboards(filePath string, userSpace string, kb *kibana.Client) error {
	if filePath == "" {
		return errors.New("You must provide filePath")
	}
	if kb == nil {
		return errors.New("You must provide kb client")
	}
	log.Debug("FilePath: ", filePath)
	log.Debug("UserSpace: ", userSpace)

	// Read file that contain all dashboard ans reference
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Import all dashboard and includes all references
	trace, err := kb.API.KibanaSavedObject.Import(
		b,
		true,
		userSpace,
	)
	if err != nil {
		return err
	}
	log.Debugf("Import trace: %+v", trace)

	return nil
}

// ImportDashboards permit to import all dashboard  and include all references
func ImportDashboards(c *cli.Context) error {
	kb, err := manageKibanaGlobalParameters(c)
	if err != nil {
		return err
	}

	importFilePath := c.String("file-path")
	userSpace := c.String("user-space")

	if importFilePath == "" {
		return errors.New("You must set --file-path")
	}
	if userSpace == "" {
		return errors.New("You must set --user-space")
	}

	err = importDashboards(importFilePath, userSpace, kb)
	if err != nil {
		return err
	}

	log.Infof("Import all dashboards successfully from %s", importFilePath)

	return nil
}
