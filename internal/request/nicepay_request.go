package request

type RegistrationRequest struct {
	Timestamp            string `json:"timeStamp"`
	MerchantId           string `json:"iMid"`
	PayMethod            string `json:"payMethod"`
	Currency             string `json:"currency"`
	Amount               int    `json:"amt"`
	ReferenceNo          string `json:"referenceNo"`
	GoodsName            string `json:"goodsNm"`
	BillingName          string `json:"billingNm"`
	BillingPhone         string `json:"billingPhone"`
	BillingEmail         string `json:"billingEmail"`
	BillingAddress       string `json:"billingAddr"`
	BillingCity          string `json:"billingCity"`
	BillingState         string `json:"billingState"`
	BillingPostalCode    string `json:"billingPostCd"`
	BillingCountry       string `json:"billingCountry"`
	DeliveryName         string `json:"deliveryNm"`
	DeliveryPhone        string `json:"deliveryPhone"`
	DeliveryAddress      string `json:"deliveryAddr"`
	DeliveryCity         string `json:"deliveryCity"`
	DeliveryState        string `json:"deliveryState"`
	DeliveryPostalCode   string `json:"deliveryPostCd"`
	DeliveryCountry      string `json:"deliveryCountry"`
	NotificationUrl      string `json:"dbProcessUrl"`
	MerchantToken        string `json:"merchantToken"`
	RequestDate          string `json:"reqDt"`
	RequestTime          string `json:"reqTm"`
	RequestDomain        string `json:"reqDomain"`
	RequestServerIP      string `json:"reqServerIP"`
	RequestClientVersion string `json:"reqClientVer"`
	UserIP               string `json:"userIP"`
	UserSessionID        string `json:"userSessionID"`
	UserAgent            string `json:"userAgent"`
	UserLanguage         string `json:"userLanguage"`
	CartData             string `json:"cartData"`
	InstallmentType      int    `json:"instmntType"`
	InstallmentMonth     int    `json:"instmntMon"`
	RecurringOption      int    `json:"recurrOpt"`
}
