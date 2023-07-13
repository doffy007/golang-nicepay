package request

type StatusRequest struct {
	Timestamp     string `json:"timeStamp"`
	TransactionId string `json:"tXid"`
	MerchantId    string `json:"iMId"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        int    `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}
