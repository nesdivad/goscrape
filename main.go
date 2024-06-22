package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Source    string
	Title     string
	Excerpt   string
	Contents  string
	CrawledAt time.Time
	LastEdit  time.Time
}

var siteflag string

func init() {
	flag.StringVar(&siteflag, "site", "", "Enter site that you would like to scrape")

	flag.Parse()
}

func main() {
	sitehost, err := parseSiteHost(siteflag)
	if err != nil {
		panic(err.Error())
	}

	c := colly.NewCollector(colly.AllowedDomains(sitehost))

	c.OnHTML("article[data-ndla-article]", func(h *colly.HTMLElement) {
		item := Item{
			Source:    h.Request.URL.String(),
			Title:     h.ChildText("h1[data-style=h1-resource]"),
			Excerpt:   h.ChildText("div[class*=ingress]"),
			Contents:  h.ChildText("p"),
			CrawledAt: time.Now(),
		}
		fmt.Fprintln(os.Stdout, []any{"Source: ", item.Source}...)
		fmt.Fprintln(os.Stdout, []any{"Crawled at: ", item.CrawledAt.Local().String()}...)
		fmt.Fprintln(os.Stdout, []any{"Title: ", item.Title}...)
		fmt.Fprintln(os.Stdout, []any{"Excerpt: ", item.Excerpt}...)
		fmt.Fprintln(os.Stdout, []any{"Contents: ", item.Contents}...)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stdout, []any{"Visiting: ", r.URL}...)
	})

	c.Visit(siteflag)
}

func parseSiteHost(site string) (string, error) {
	url, err := url.Parse(site)
	if err != nil {
		return "", err
	}

	return url.Host, nil
}
