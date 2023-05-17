package elktools_kibana

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/disaster37/elktools/v8/helper"
	"github.com/disaster37/go-kibana-rest/v8"
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
	b, err := os.ReadFile(filePath)
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

func MoveDataView(c *cli.Context) error {
	kb, err := manageKibanaGlobalParameters(c)
	if err != nil {
		return err
	}

	currentDataView := c.String("current-data-view")
	targetDataView := c.String("target-data-view")
	userSpace := c.String("user-space")

	if currentDataView == "" {
		return errors.New("You must set --current-data-view")
	}
	if targetDataView == "" {
		return errors.New("You must set --target-data-view")
	}
	if userSpace == "" {
		return errors.New("You must set --user-space")
	}

	err = moveDataView(userSpace, currentDataView, targetDataView, kb)
	if err != nil {
		return err
	}

	log.Infof("Move all resources in '%s' from '%s' to '%s' data view", userSpace, currentDataView, targetDataView)

	return nil
}

func moveDataView(userSpace string, currentDataView string, targetDataView string, kb *kibana.Client) error {

	if currentDataView == "" {
		return errors.New("You must provide currentDataView")
	}
	if targetDataView == "" {
		return errors.New("You must provide targetDataView")
	}
	if kb == nil {
		return errors.New("You must provide kb client")
	}

	if userSpace == "" {
		userSpace = "default"
	}
	log.Debug("currentDataView: ", currentDataView)
	log.Debug("targetDataView: ", targetDataView)
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

	// Split json
	jsonStrs := strings.Split(string(data), "\n")

	for iJson, jsonStr := range jsonStrs {
		var j map[string]any
		var expectedJ []byte

		if err := json.Unmarshal([]byte(jsonStr), &j); err != nil {
			return err
		}

		if j["type"] != "dashboard" && j["type"] != "visualization" && j["type"] != "search" && j["type"] != "lens" && j["type"] != "map" {
			log.Infof("Skip object %s, because '%s'", j["id"], j["type"])
			continue
		}

		refs := j["references"].([]any)
		for iRef, ref := range refs {
			refTmp := ref.(map[string]any)
			if refTmp["id"] == currentDataView && refTmp["type"] == "index-pattern" {
				refTmp["id"] = targetDataView
			}
			refs[iRef] = refTmp
		}

		expectedJ, err = json.Marshal(j)
		if err != nil {
			return err
		}

		jsonStrs[iJson] = string(expectedJ)
	}

	res := strings.Join(jsonStrs, "\n")
		trace, err := kb.API.KibanaSavedObject.Import([]byte(res), true, userSpace)
		if err != nil {
			return err
		}

		log.Debugf("Import trace: %s", spew.Sprint(trace))

	return nil
}
