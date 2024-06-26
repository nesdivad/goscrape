package structs

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Item struct {
	Source    string
	Title     string
	Excerpt   string
	Contents  string
	CrawledAt time.Time
}

func ToItem(h *colly.HTMLElement, rule Rule) Item {
	return Item{
		Source:    h.Request.URL.String(),
		Title:     h.ChildText(rule.TitleSelector),
		Excerpt:   h.ChildText(rule.ExcerptSelector),
		Contents:  h.ChildText(rule.ContentSelector),
		CrawledAt: time.Now(),
	}
}

func (item Item) String() string {
	return fmt.Sprintf(
		"Source: %s \n Crawled at: %s \n Title: %s \n Excerpt: %s \n Contents: %s \n\n",
		item.Source,
		item.CrawledAt.UTC().String(),
		item.Title,
		item.Excerpt,
		item.Contents,
	)
}
