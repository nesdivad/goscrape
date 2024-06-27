package structs

import (
	"errors"
)

type Rule struct {
	QuerySelector   string `json:"querySelector"`
	TitleSelector   string `json:"titleSelector"`
	ExcerptSelector string `json:"excerptSelector"`
	ContentSelector string `json:"contentSelector"`
}

func (r *Rule) Validate() error {
	if r.QuerySelector == "" {
		return errors.New("Rule is missing QuerySelector")
	}
	if r.TitleSelector == "" {
		return errors.New("Rule is missing TitleSelector")
	}
	if r.ContentSelector == "" {
		return errors.New("Rule is missing ContentSelector")
	}
	return nil
}
