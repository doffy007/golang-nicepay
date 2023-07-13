package response

type RegistrationResponse struct {
	ResultCode              int    `json:"resultCd"`
	ResultMessage           string `json:"resultMsg"`
	TransactionID           string `json:"tXid"`
	ReferenceNo             string `json:"referenceNo"`
	PayMethod               string `json:"payMethod"`
	Amount                  string `json:"amt"`
	TransactionDate         string `json:"transDt"`
	TransactionTime         string `json:"transTm"`
	Description             string `json:"description"`
	Bank                    string `json:"bankCd"`
	VirtualAccountNumber    string `json:"vacctNo"`
	Mitra                   string `json:"mitraCd"`
	PayNumber               string `json:"payNo"`
	Currency                string `json:"currency"`
	GoodsName               string `json:"goodsNm"`
	BillingName             string `json:"billingNm"`
	VirtualAccountValidDate string `json:"vacctValidDt"`
	VirtualAccountValidTime string `json:"vacctValidTm"`
	PayValidDate            string `json:"payValidDt"`
	PayValidTime            string `json:"payValidTm"`
	RequestUrl              string `json:"requestURL"`
	PaymentExpDate          string `json:"paymentExpDt"`
	PaymentExpTime          string `json:"paymentExpTm"`
	QrContent               string `json:"qrContent"`
	QrUrl                   string `json:"qrUrl"`
}
