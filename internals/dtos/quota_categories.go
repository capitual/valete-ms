package dtos

type QuotaCategoryDto struct {
	Name   string `json:"name"`
	Spread string `json:"spread"`
	Ttls   int    `json:"ttls"`
}
