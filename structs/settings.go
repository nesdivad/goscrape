package structs

type Settings struct {
	Depth      int         `json:"depth"`
	LimitRules []LimitRule `json:"limitRules"`
}
