package mofid

type SearchSymbolReq struct {
	Term string `json:"Term"`
}

type SearchSymbolRes []struct {
	Isin           string `json:"Isin"`
	Label          string `json:"Label"`
	MarketUnitName string `json:"MarketUnitName"`
	State          string `json:"State"`
	Value          int    `json:"Value"`
}
