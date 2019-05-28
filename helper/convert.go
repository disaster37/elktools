package helper

import(
    "encoding/json"
    log "github.com/sirupsen/logrus"
)


// JsonToBytes permit to convert interface that represent Json to bytes
func JsonToBytes(data interface{}) ([]byte, error) {
    
    log.Debugf("Data: %s", data)
    
    dataBytes, err := json.MarshalIndent(data,"", "  ")
    if err != nil {
        return nil, err
    }
    
    return dataBytes, nil
}

// BytesToJson permit to convert bytes to Json struct
func BytesToJson(data []byte, convertedData interface{}) (error) {
    log.Debugf("Data: %s", data)
    
    err := json.Unmarshal(data, convertedData)
    
    if err != nil {
        return err
    }
    
    return nil
}