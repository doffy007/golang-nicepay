package entity

import "time"

type Status struct {
	ID               int    `json:"id" db:"id"`
	Timestamp        string `json:"timeStamp" db:"timeStamp"`
	TransactionId    string `json:"tXid" db:"tXid"`
	MerchantId       string `json:"iMid" db:"iMid"`
	Currency         string `json:"currency" db:"currency"`
	Amount           string `json:"amt" db:"amount"`
	InstallmentMonth int    `json:"instmntMon" db:"instmntMon"`
	InstallmentType  int    `json:"instmntType" db:"instmntType"`
	ReferenceNo      string `json:"referenceNo" db:"referenceNo"`
	GoodsName        string `json:"goodsNm" db:"goodsNm"`
	PayMethod        string `json:"payMethod" db:"payMethod"`
	BillingName      string `json:"billingNm" db:"billingNm"`
	RequestDate      string `json:"reqDt" db:"reqDt"`
	RequestTime      string `json:"reqTm" db:"reqTm"`
	Status           string `json:"status" db:"status"`
	ResultCode       string `json:"resultCd" db:"resultCd"`
	ResultMessage    string `json:"resultMsg" db:"resultMsg"`

	CardNo       string `json:"cardNo" db:"cardNo"`
	PreauthToken string `json:"preauthToken" db:"preauthToken"`
	AcquBankCode string `json:"acquBankCd" db:"acquBankCd"`
	IssuBankCode string `json:"issuBankCd" db:"issuBankCd"`

	VacctValidDt string `json:"vacctValidDt" db:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm" db:"vacctValidTm"`
	VacctNo      string `json:"vacctNo" db:"vacctNo"`
	BankCd       string `json:"bankCd" db:"bankCd"`

	PayNo       string `json:"payNo" db:"payNo"`
	MitraCode   string `json:"mitraCd" db:"mitraCd"`
	ReceiptCode string `json:"receiptCode" db:"receiptCode"`
	PayValidDt  string `json:"payValidDt" db:"payValidDt"`
	PayValidTm  string `json:"payValidTm" db:"payValidTm"`

	MRefNumber           string `json:"mRefNo" db:"mRefNo"`
	AcquStatus           string `json:"acquStatus" db:"acquStatus"`
	CardExpYymm          string `json:"cardExpYymm" db:"cardExpYymm"`
	AcquBankName         string `json:"acquBankNm" db:"acquBankNm"`
	IssuBankName         string `json:"issuBankNm" db:"issuBankNm"`
	DepositDate          string `json:"depositDt" db:"depositDt"`
	DepositTime          string `json:"depositTm" db:"depositTm"`
	PaymentExpDate       string `json:"paymentExpDt" db:"paymentExpDt"`
	PaymentExpTime       string `json:"paymentExpTm" db:"paymentExpTm"`
	PaymentTransactionSn string `json:"paymentTrxSn" db:"paymentTrxSn"`
	CancelTransactionSn  string `json:"cancelTrxSn" db:"cancelTrxSn"`
	UserId               string `json:"userId" db:"userId"`
	ShopId               string `json:"shopId" db:"shopId"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
