package structs

import "errors"

type Output struct {
	Path     string `json:"path"`
	Filetype string `json:"fileType"`
	Filename string `json:"fileName"`
}

func (o *Output) Validate() error {
	if o.Path == "" {
		return errors.New("missing field 'path'")
	}
	if o.Filetype == "" {
		o.Filetype = "json"
	}
	if o.Filename == "" {
		o.Filename = "scrape"
	}
	return nil
}
