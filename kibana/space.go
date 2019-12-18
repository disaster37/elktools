package elktools_kibana

import (
	"fmt"

	"github.com/disaster37/elktools/v7/helper"
	"github.com/disaster37/go-kibana-rest/v7"
	"github.com/disaster37/go-kibana-rest/v7/kbapi"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type SpaceSetting struct {
	Name           string   `yaml:"name"`
	Description    string   `yaml:"description,omitempty"`
	Features       []string `yaml:"features"`
	LoadDashboards bool     `yaml:"load_dashboards"`
}

func createSpaceFromFile(filePath string, kb *kibana.Client) error {
	if filePath == "" {
		return errors.New("You must provide file")
	}

	if kb == nil {
		return errors.New("You must provide kb client")
	}

	log.Debug("File: ", filePath)

	// Load Yaml spec to ccreate create space
	spaceSetting := &SpaceSetting{}
	err := helper.LoadYaml(filePath, spaceSetting)
	if err != nil {
		return err
	}

	// Manage space
	space, err := kb.API.KibanaSpaces.Get(spaceSetting.Name)
	if err != nil {
		return err
	}
	// Check if space already exist
	if space != nil {
		// Update space
		space.Description = spaceSetting.Description
		space, err = kb.API.KibanaSpaces.Update(space)
		if err != nil {
			return err
		}
		space.Description = spaceSetting.Description
		space, err = kb.API.KibanaSpaces.Update(space)
		if err != nil {
			return nil
		}
		log.Infof("Space %s updated successfully", spaceSetting.Name)
	} else {
		// Create new space
		space := &kbapi.KibanaSpace{
			ID:          spaceSetting.Name,
			Name:        spaceSetting.Name,
			Description: spaceSetting.Description,
		}
		space, err = kb.API.KibanaSpaces.Create(space)
		if err != nil {
			return nil
		}
		log.Infof("Space %s created successsfully", spaceSetting.Name)
	}

	// Manage roles
	roleName := fmt.Sprintf("space_%s", spaceSetting.Name)
	role, err := kb.API.KibanaRoleManagement.Get(roleName)
	if err != nil {
		return err
	}
	var state string
	if role != nil {
		// Update role
		state = "updated"
	} else {
		// Create new role
		state = "created"
	}
	role = &kbapi.KibanaRole{
		Name: roleName,
		Kibana: []kbapi.KibanaRoleKibana{
			{
				Feature: map[string][]string{},
				Spaces: []string{
					spaceSetting.Name,
				},
			},
		},
	}
	for _, feature := range spaceSetting.Features {
		role.Kibana[0].Feature[feature] = []string{"all"}
	}
	role, err = kb.API.KibanaRoleManagement.CreateOrUpdate(role)
	if err != nil {
		return err
	}
	log.Infof("Role %s %s successsfully", roleName, state)

	// Manage dashboards
	if spaceSetting.LoadDashboards == true {
		// Get all TPL dashboards
		parameters := &kbapi.OptionalFindParameters{
			Fields:         []string{"id", "title"},
			ObjectsPerPage: 10000,
			SearchFields:   []string{"title"},
			Search:         "TPL*",
		}
		dashboardsRes, err := kb.API.KibanaSavedObject.Find("dashboard", "default", parameters)
		if err != nil {
			return err
		}
		dashboards := dashboardsRes["saved_objects"].([]interface{})
		for _, dashboard := range dashboards {
			dashboardTemp := dashboard.(map[string]interface{})
			log.Debugf("Process dashboard %s (%s)", dashboardTemp["id"].(string), dashboardTemp["title"].(string))
		}
	}

	return nil
}
