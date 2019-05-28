package helper

import(
    "os"
    log "github.com/sirupsen/logrus"
    "github.com/pkg/errors"
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

