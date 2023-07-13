package response

type PaymentResponse struct {
	ResultCode      int         `json:"resultCd"`
	ResultMessage   string      `json:"resultMsg"`
	TransactionID   string      `json:"tXid"`
	ReferenceNo     string      `json:"referenceNo"`
	PaymentMethod   string      `json:"payMethod"`
	Amount          string      `json:"amt"`
	CancelAmount    string      `json:"cancelAmt"`
	RequestDate     string      `json:"reqDt"`
	RequestTime     string      `json:"reqTm"`
	TransactionDate string      `json:"transDt"`
	TransactionTime string      `json:"transTm"`
	DepositDate     string      `json:"depositDt"`
	DepositTime     string      `json:"depositTm"`
	Currency        string      `json:"currency"`
	GoodsName       string      `json:"goodsNm"`
	BillingName     string      `json:"billingNm"`
	Status          string      `json:"status"`
	Data            interface{} `json:"data"`
}
