package structs

type Rule struct {
	QuerySelector   string `json:"querySelector"`
	TitleSelector   string `json:"titleSelector"`
	ExcerptSelector string `json:"excerptSelector"`
	ContentSelector string `json:"contentSelector"`
}
