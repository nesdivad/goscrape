package utils

import (
	"bytes"
	"fmt"
	"goscrape/structs"
	"os"
)

func WriteJsonl(items []structs.Item, path string) error {
	var w bytes.Buffer
	for i, item := range items {
		buffer := new(bytes.Buffer)
		marshal, err := Marshal(item)
		if err != nil {
			return err
		}
		if err := Compact(buffer, marshal); err != nil {
			return err
		}
		if i+1 == len(items) {
			_, err = fmt.Fprint(&w, buffer)
		} else {
			_, err = fmt.Fprintln(&w, buffer)
		}
		if err != nil {
			return fmt.Errorf("could not write result to buffer.\nErrors: %s", err)
		}
	}

	if err := os.WriteFile(path, w.Bytes(), 0644); err != nil {
		return fmt.Errorf("could not write to file.\nErrors: %s", err)
	}

	return nil
}
