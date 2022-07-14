package dtos

type QuoteCategoryDto struct {
	Name   string  `json:"name"`
	Spread float64 `json:"spread"`
	Ttls   string  `json:"ttls"`
}
