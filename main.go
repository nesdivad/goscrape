package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goscrape/structs"
	"net/url"

	"os"

	"github.com/gocolly/colly"
)

var configflag string

func init() {
	flag.StringVar(&configflag, "config", "", "Path to config file")
	flag.Parse()
}

func main() {
	config, err := parseConfig(configflag)
	if err != nil {
		panic(fmt.Sprintf("Could not parse configuration file. \n Error: %s", err))
	}

	sitehost, err := parseSiteHost(config.URL)
	if err != nil {
		panic(fmt.Sprintf("Could not parse url. \n Error: %s", err))
	}

	c := colly.NewCollector(
		colly.AllowedDomains(sitehost),
		colly.MaxDepth(config.Depth),
	)
	c.AllowURLRevisit = false
	c.DisallowedURLFilters = structs.GetRegex(config.URLFilters)

	for _, filter := range config.URLFilters {
		fmt.Fprintln(os.Stdout, []any{"Filter: ", &filter.Regexp}...)
	}

	for _, rule := range config.Rules {
		c.OnHTML(rule.QuerySelector, func(h *colly.HTMLElement) {
			item := structs.ToItem(h, rule)
			fmt.Println(item.String())
		})
	}

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		c.Visit(h.Attr("href"))
	})

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

func parseConfig(filename string) (*structs.Config, error) {
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
