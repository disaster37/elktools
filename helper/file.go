package helper

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// WriteFile permit to write the content in file
// It return error if something wrong when write file
func WriteFile(path string, content []byte) error {

	if path == "" {
		return errors.New("Path can't be empty")
	}
	if len(content) == 0 {
		return errors.New("Content can't be empty")
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	log.Debugf("Write file %s successfully with content: %s", path, content)

	return nil
}

// ListFilesInPath permtit to list all file in provided path that match the extension
// It return error if somethink wrong when it list the file on path
// It return a list of full path.
// It return empty list if it doesn't found file
func ListFilesInPath(path string, extention string) ([]string, error) {
	if path == "" {
		return nil, errors.New("Path can't be empty")
	}

	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == extention {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// LoadYaml permit to read yaml file and convert content to map
func LoadYaml(filePath string, data interface{}) error {
	if filePath == "" {
		return errors.New("FilePath can't be empty")
	}

	log.Debug("FilePath: ", filePath)

	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, data)
	if err != nil {
		return err
	}

	return nil
}
