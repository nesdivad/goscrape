package main

import (
	"flag"
	"fmt"
	"goscrape/statistics"
	"goscrape/structs"
	"goscrape/utils"
	"time"

	"github.com/gocolly/colly"
)

var configflag string
var configjsonflag string
var verbose bool
var stats *statistics.Statistics

func init() {
	flag.BoolVar(&verbose, "v", false, "Use flag if you want verbose logs.")
	flag.StringVar(&configflag, "config", "", "Path to config file")
	flag.StringVar(&configjsonflag, "configjson", "", "Config as a json-string. Compact version works best.")
	flag.Parse()

	if verbose {
		stats = statistics.New()
	}
}

func main() {
	config, err := utils.ParseConfig(configflag, configjsonflag)
	if err != nil {
		panic(fmt.Sprintf("Could not parse configuration file.\nError: %s", err))
	}
	if err = config.Validate(); err != nil {
		panic(fmt.Sprintf("Config validation failed.\nError: %s", err))
	}

	sitehost, err := utils.ParseSiteHost(config.URL)
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
		if verbose {
			fmt.Printf("Visiting: %s\n", r.URL)
			stats.NumberOfPages += 1
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while visiting: %s\t Status code: %d\nErrors: %s\n", r.Request.URL, r.StatusCode, err)
	})

	c.Visit(config.URL)
	c.Wait()

	if config.Output.Filetype == "json" {
		for _, item := range items {
			path := fmt.Sprintf("%s/%s.json", config.Output.Path, item.Title)
			bytesWritten, err := utils.WriteJson(item, path)
			if err != nil {
				panic(err)
			}
			if verbose {
				stats.BytesWritten += bytesWritten
			}
		}
	} else if config.Output.Filetype == "jsonl" {
		bytesWritten, err := utils.WriteJsonl(items, config.Output)
		if err != nil {
			panic(err)
		}
		if verbose {
			stats.BytesWritten = bytesWritten
		}
	}

	if verbose {
		stats.Print()
	}
}
