package helper

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Permit to write file
func WriteFile(path string, content []byte) error {

	if path == "" {
		errors.New("Path can't be empty")
	}
	if len(content) == 0 {
		errors.New("Content can't be empty")
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

func ListFilesInPath(path string, extention string) ([]string, error) {
	if path == "" {
		errors.New("Path can't be empty")
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
