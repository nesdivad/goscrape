package structs

import "errors"

type Config struct {
	Rules      []Rule      `json:"rules"`
	URL        string      `json:"url"`
	Settings   Settings    `json:"settings"`
	URLFilters []URLFilter `json:"urlFilters"`
	Output     Output      `json:"output"`
}

func (c *Config) Validate() error {
	for _, r := range c.Rules {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	if c.URL == "" {
		return errors.New("missing or empty field 'URL'")
	}

	for _, u := range c.URLFilters {
		if err := u.Validate(); err != nil {
			return err
		}
	}

	if err := c.Output.Validate(); err != nil {
		return err
	}

	return nil
}
