package request

type PaymentRequest struct {
	Timestamp      string `json:"timeStamp"`
	TransactionId  string `json:"tXid"`
	CCNumber       string `jsong:"cardNo"`
	CardExpiry     string `json:"cardExpYymm"`
	CardCVV        string `json:"cardCvv"`
	CardHolderName string `json:"cardHolderNm"`
	RecurringToken string `json:"recurringToken"`
	PreAuthToken   string `json:"preauthToken"`
	MerchantToken  string `json:"merchantToken"`
	CallbackUrl    string `json:"callBackUrl"`
}
