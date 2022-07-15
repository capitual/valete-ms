package dtos

type QuoteCategoryDto struct {
	Name   string `json:"name"`
	Spread string `json:"spread"`
	Ttls   int64  `json:"ttls"`
}
