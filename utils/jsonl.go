package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goscrape/structs"
	"os"
)

func WriteJsonl(items []structs.Item, output structs.Output) (int, error) {
	chunks := chunkBy(items, output.Chunk)
	bytesWritten := 0

	for i, chunk := range chunks {
		var w bytes.Buffer
		for j, item := range chunk {
			buffer := new(bytes.Buffer)
			marshal, err := marshal(item)
			if err != nil {
				return bytesWritten, err
			}
			if err := compact(buffer, marshal); err != nil {
				return bytesWritten, err
			}
			if j+1 == len(chunk) {
				_, err = fmt.Fprint(&w, buffer)
			} else {
				_, err = fmt.Fprintln(&w, buffer)
			}
			if err != nil {
				return bytesWritten, fmt.Errorf("could not write result to buffer.\nErrors: %s", err)
			}
		}

		path := fmt.Sprintf("%s/%s_%d.jsonl", output.Path, output.Filename, i+1)
		if err := os.WriteFile(path, w.Bytes(), 0644); err != nil {
			return bytesWritten, fmt.Errorf("could not write to file.\nErrors: %s", err)
		}

		bytesWritten += binary.Size(w.Bytes())
	}

	return bytesWritten, nil
}
