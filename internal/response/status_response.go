package response

type StatusResponse struct {
	TransactionId    string `json:"tXid"`
	MerchantId       string `json:"iMid"`
	Currency         string `json:"currency"`
	Amount           string `json:"amt"`
	InstallmentMonth int    `json:"instmntMon"`
	InstallmentType  int    `json:"instmntType"`
	ReferenceNo      string `json:"referenceNo"`
	GoodsName        string `json:"goodsNm"`
	PayMethod        string `json:"payMethod"`
	BillingName      string `json:"billingNm"`
	RequestDate      string `json:"reqDt"`
	RequestTime      string `json:"reqTm"`
	Status           string `json:"status"`
	ResultCode       int    `json:"resultCd"`
	ResultMessage    string `json:"resultMsg"`

	CardNo       string `json:"cardNo"`
	PreauthToken string `json:"preauthToken"`
	AcquBankCode string `json:"acquBankCd"`
	IssuBankCode string `json:"issuBankCd"`

	VacctValidDt string `json:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm"`
	VacctNo      string `json:"vacctNo"`
	BankCd       string `json:"bankCd"`

	PayNo       string `json:"payNo"`
	MitraCode   string `json:"mitraCd"`
	ReceiptCode string `json:"receiptCode"`
	PayValidDt  string `json:"payValidDt"`
	PayValidTm  string `json:"payValidTm"`

	MRefNumber           string `json:"mRefNo"`
	AcquStatus           string `json:"acquStatus"`
	CardExpYymm          string `json:"cardExpYymm"`
	AcquBankName         string `json:"acquBankNm"`
	IssuBankName         string `json:"issuBankNm"`
	DepositDate          string `json:"depositDt"`
	DepositTime          string `json:"depositTm"`
	PaymentExpDate       string `json:"paymentExpDt"`
	PaymentExpTime       string `json:"paymentExpTm"`
	PaymentTransactionSn string `json:"paymentTrxSn"`
	CancelTransactionSn  string `json:"cancelTrxSn"`
	UserId               string `json:"userId"`
	ShopId               string `json:"shopId"`
}
