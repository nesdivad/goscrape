package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goscrape/structs"
	"net/url"
	"time"

	"os"

	"github.com/gocolly/colly"
)

func main() {
	var configflag string
	flag.StringVar(&configflag, "config", "", "Path to config file")
	flag.Parse()

	config, err := parseFile(configflag)
	if err != nil {
		panic(fmt.Sprintf("Could not parse configuration file. \n Error: %s", err))
	}

	sitehost, err := parseSiteHost(config.URL)
	if err != nil {
		panic(err.Error())
	}

	c := colly.NewCollector(colly.AllowedDomains(sitehost))

	for _, rule := range config.Rules {
		c.OnHTML(rule.QuerySelector, func(h *colly.HTMLElement) {
			item := structs.Item{
				Source:    h.Request.URL.String(),
				Title:     h.ChildText(rule.TitleSelector),
				Excerpt:   h.ChildText(rule.ExcerptSelector),
				Contents:  h.ChildText(rule.ContentSelector),
				CrawledAt: time.Now(),
			}
			fmt.Fprintln(os.Stdout, []any{"Source: ", item.Source}...)
			fmt.Fprintln(os.Stdout, []any{"Crawled at: ", item.CrawledAt.Local().String()}...)
			fmt.Fprintln(os.Stdout, []any{"Title: ", item.Title}...)
			fmt.Fprintln(os.Stdout, []any{"Excerpt: ", item.Excerpt}...)
			fmt.Fprintln(os.Stdout, []any{"Contents: ", item.Contents}...)
		})
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stdout, []any{"Visiting: ", r.URL}...)
	})

	c.Visit(config.URL)
}

func parseSiteHost(site string) (string, error) {
	url, err := url.Parse(site)
	if err != nil {
		return "", err
	}

	return url.Host, nil
}

func parseFile(filename string) (*structs.Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := structs.Config{}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
