package structs

type LimitRule struct {
	DomainGlob  string `json:"domainGlob"`
	Parallelism int    `json:"parallelism"`
	RandomDelay int    `json:"randomDelay"`
}
