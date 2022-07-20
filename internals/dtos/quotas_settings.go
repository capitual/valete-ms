package dtos

type QuoteSettingDto struct {
	Currency  uint `json:"currency"`
	ExpiresIn int  `json:"expires_in"`
	PartnerId uint `json:"partner_id"`
}
