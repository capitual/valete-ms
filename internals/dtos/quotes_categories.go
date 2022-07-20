package dtos

type QuoteCategoryDto struct {
	Name      string `json:"name"`
	SpreadAsk string `json:"spread_ask"`
	SpreadBid string `json:"spread_bid"`
	Ttls      int    `json:"ttls"`
}
