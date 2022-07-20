package dtos

type QuotePartnerDataDto struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	QuotePartnerID uint   `json:"quote_partner_id"`
	Environment    string `json:"environment"`
}
