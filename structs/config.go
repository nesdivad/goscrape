package structs

type Config struct {
	Rules []Rule `json:"rules"`
	URL   string `json:"url"`
	Depth int    `json:"depth"`
}
