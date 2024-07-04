package statistics

import (
	"encoding/json"
	"fmt"
)

type Statistics struct {
	BytesWritten  int `json:"bytesWritten"`
	NumberOfPages int `json:"numberOfPages"`
	//TimeTaken     time.Timer `json:"timeTaken"`
	//NumberOfNoResults int        `json:"numberOfNoResults"`
}

func New() *Statistics {
	return &Statistics{
		BytesWritten:  0,
		NumberOfPages: 0,
	}
}

func (s *Statistics) Print() {
	j, _ := json.MarshalIndent(&s, "", "  ")
	fmt.Println("=============================================")
	fmt.Printf("Statistics:\n%s\n", j)
	fmt.Println("=============================================")
}
