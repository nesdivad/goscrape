package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goscrape/structs"
	"os"
)

func Marshal(item structs.Item) ([]byte, error) {
	marshal, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("could not marshal response.\nErrors: %s", err)
	}
	return marshal, nil
}

func Compact(buffer *bytes.Buffer, bytes []byte) error {
	if err := json.Compact(buffer, bytes); err != nil {
		return fmt.Errorf("could not compact json.\nErrors: %s", err)
	}
	return nil
}

func WriteJson(item structs.Item, path string) error {
	marshal, err := Marshal(item)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, marshal, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file.\nErrors: %s", err)
	}

	return nil
}
