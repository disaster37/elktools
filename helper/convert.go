package helper

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// JsonToBytes permit to convert interface that represent Json to bytes
// The json is pretty formatted before to be converted in bytes.
// It return error if something wrong when convert json in byte
// It return an array of bytes that represent the pretty formated json
func JsonToBytes(data interface{}) ([]byte, error) {

	log.Debugf("Data: %s", data)

	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

// BytesToJson permit to convert bytes to Json struct
// It return error if somthing wrong when it convert json to struct
// It add struct on convertedData interface.
func BytesToJson(data []byte, convertedData interface{}) error {
	log.Debugf("Data: %s", data)

	err := json.Unmarshal(data, convertedData)

	if err != nil {
		return err
	}

	return nil
}
