package structs

import "time"

type Item struct {
	Source    string
	Title     string
	Excerpt   string
	Contents  string
	CrawledAt time.Time
}
