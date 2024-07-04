package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goscrape/structs"
	"os"
)

func marshal(item structs.Item) ([]byte, error) {
	marshal, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("could not marshal response.\nErrors: %s", err)
	}
	return marshal, nil
}

func compact(buffer *bytes.Buffer, bytes []byte) error {
	if err := json.Compact(buffer, bytes); err != nil {
		return fmt.Errorf("could not compact json.\nErrors: %s", err)
	}
	return nil
}

func WriteJson(item structs.Item, path string) (int, error) {
	marshal, err := marshal(item)
	if err != nil {
		return 0, err
	}
	err = os.WriteFile(path, marshal, 0644)
	if err != nil {
		return 0, fmt.Errorf("could not write to file.\nErrors: %s", err)
	}

	return binary.Size(marshal), nil
}
