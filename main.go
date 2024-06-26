package main

import (
	"bytes"
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
		colly.Async(true),
	)
	c.AllowURLRevisit = false
	c.DisallowedURLFilters = structs.GetRegex(config.URLFilters)

	items := []structs.Item{}

	for _, rule := range config.Rules {
		c.OnHTML(rule.QuerySelector, func(h *colly.HTMLElement) {
			item := structs.ToItem(h, rule)
			items = append(items, item)
		})
	}

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		c.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintln(os.Stdout, []any{"Visiting: ", r.URL}...)
	})

	c.Visit(config.URL)
	c.Wait()

	if config.Output.Filetype == "json" {
		for _, item := range items {
			marshal, err := json.Marshal(item)
			if err != nil {
				panic(fmt.Sprintf("Could not marshal response. \n Errors: %s", err))
			}
			err = os.WriteFile(fmt.Sprintf("%s/%s.json", config.Output.Path, item.Title), marshal, 0644)
			if err != nil {
				panic(fmt.Sprintf("Could not write to file. \n Errors: %s", err))
			}
		}
	} else if config.Output.Filetype == "jsonl" {
		var w bytes.Buffer
		for _, item := range items {
			buffer := new(bytes.Buffer)
			marshal, err := json.Marshal(item)
			if err != nil {
				panic(fmt.Sprintf("Could not marshal response. \n Errors: %s", err))
			}
			err = json.Compact(buffer, marshal)
			if err != nil {
				panic(fmt.Sprintf("Could not compact json. \n Errors: %s", err))
			}
			_, err = fmt.Fprintln(&w, buffer)
			if err != nil {
				panic(fmt.Sprintf("Could not write result to buffer. \n Errors: %s", err))
			}
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s.jsonl", config.Output.Path, config.Output.Filename), w.Bytes(), 0644)
		if err != nil {
			panic(fmt.Sprintf("Could not write to file. \n Errors: %s", err))
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
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
