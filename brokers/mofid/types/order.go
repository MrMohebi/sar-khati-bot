package mofid

type OrderReq struct {
	IsSymbolCautionAgreement  bool   `json:"IsSymbolCautionAgreement"`
	CautionAgreementSelected  bool   `json:"CautionAgreementSelected"`
	IsSymbolSepahAgreement    bool   `json:"IsSymbolSepahAgreement"`
	SepahAgreementSelected    bool   `json:"SepahAgreementSelected"`
	OrderCount                int    `json:"orderCount"`
	OrderPrice                int    `json:"orderPrice"`
	FinancialProviderID       int    `json:"FinancialProviderId"`
	MinimumQuantity           int    `json:"minimumQuantity"`
	MaxShow                   int    `json:"maxShow"`
	OrderID                   int    `json:"orderId"`
	Isin                      string `json:"isin"`
	OrderSide                 int    `json:"orderSide"`
	OrderValidity             int    `json:"orderValidity"`
	OrderValiditydate         any    `json:"orderValiditydate"`
	ShortSellIsEnabled        bool   `json:"shortSellIsEnabled"`
	ShortSellIncentivePercent int    `json:"shortSellIncentivePercent"`
}

type SendOrderRes struct {
	Data          any    `json:"Data"`
	MessageDesc   string `json:"MessageDesc"`
	IsSuccessfull bool   `json:"IsSuccessfull"`
	MessageCode   any    `json:"MessageCode"`
	Version       int    `json:"Version"`
}
