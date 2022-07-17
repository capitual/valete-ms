package dtos

type QuotaSettingDto struct {
	Currency  string `json:"currency"`
	ExpiresIn int    `json:"expires_in"`
	PartnerId uint   `json:"partner_id"`
}
