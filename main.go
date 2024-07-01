package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goscrape/structs"
	"goscrape/utils"
	"net/url"
	"time"

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
		panic(fmt.Sprintf("Could not parse configuration file.\nError: %s", err))
	}
	if err = config.Validate(); err != nil {
		panic(fmt.Sprintf("Config validation failed.\nError: %s", err))
	}

	sitehost, err := parseSiteHost(config.URL)
	if err != nil {
		panic(fmt.Sprintf("Could not parse url.\nError: %s", err))
	}

	c := colly.NewCollector(
		colly.AllowedDomains(sitehost),
		colly.Async(true),
	)
	if config.Settings.Depth > 0 {
		c.MaxDepth = config.Settings.Depth
	}
	c.AllowURLRevisit = false
	c.DisallowedURLFilters = structs.GetRegex(config.URLFilters)

	for _, limit := range config.Settings.LimitRules {
		c.Limit(&colly.LimitRule{
			DomainGlob:  limit.DomainGlob,
			RandomDelay: time.Duration(limit.RandomDelay) * time.Second,
			Parallelism: limit.Parallelism,
		})
	}

	items := []structs.Item{}

	for _, rule := range config.Rules {
		c.OnHTML(rule.QuerySelector, func(h *colly.HTMLElement) {
			item := structs.ToItem(h, rule)
			if item.Contents != "" {
				items = append(items, item)
			}
		})
	}

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		c.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting: %s\t Status code: %d\nErrors: %s\n", r.Request.URL, r.StatusCode, err)
	})

	c.Visit(config.URL)
	c.Wait()

	if config.Output.Filetype == "json" {
		for _, item := range items {
			path := fmt.Sprintf("%s/%s.json", config.Output.Path, item.Title)
			if err := utils.WriteJson(item, path); err != nil {
				panic(err)
			}
		}
	} else if config.Output.Filetype == "jsonl" {
		path := fmt.Sprintf("%s/%s.jsonl", config.Output.Path, config.Output.Filename)
		if err := utils.WriteJsonl(items, path); err != nil {
			panic(err)
		}
	}
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
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
