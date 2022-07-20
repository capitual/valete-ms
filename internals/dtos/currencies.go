package dtos

type CurrencyDto struct {
	Name    string `json:"name"`
	IsoCode string `json:"iso_code"`
	Type    string `json:"type"`
}
