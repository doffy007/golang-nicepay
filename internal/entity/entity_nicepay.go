package entity

import "time"

type Nicepay struct {
	ID                   int       `json:"id" db:"id"`
	Timestamp            string    `json:"timeStamp" db:"timeStamp"`
	MerchantId           string    `json:"iMid" db:"iMid"`
	PayMethod            string    `json:"payMethod" db:"payMethod"`
	Currency             string    `json:"currency" db:"currency"`
	Amount               int       `json:"amt" db:"amt"`
	ReferenceNo          string    `json:"referenceNo" db:"referenceNo"`
	GoodsName            string    `json:"goodsNm" db:"goodsNm"`
	BillingName          string    `json:"billingNm" db:"billingNm"`
	BillingPhone         string    `json:"billingPhone" db:"billingPhone"`
	BillingEmail         string    `json:"billingEmail" db:"billingEmail"`
	BillingAddress       string    `json:"billingAddr" db:"billingAddr"`
	BillingCity          string    `json:"billingCity" db:"billingCity"`
	BillingState         string    `json:"billingState" db:"billingState"`
	BillingPostalCode    string    `json:"billingPostCd" db:"billingPostCd"`
	BillingCountry       string    `json:"billingCountry" db:"billingCountry"`
	DeliveryName         string    `json:"deliveryNm" db:"deliveryNm"`
	DeliveryPhone        string    `json:"deliveryPhone" db:"deliveryPhone"`
	DeliveryAddress      string    `json:"deliveryAddr" db:"deliveryAddr"`
	DeliveryCity         string    `json:"deliveryCity" db:"deliveryCity"`
	DeliveryState        string    `json:"deliveryState" db:"deliveryState"`
	DeliveryPostalCode   string    `json:"deliveryPostCd" db:"deliveryPostCd"`
	DeliveryCountry      string    `json:"deliveryCountry" db:"deliveryCountry"`
	NotificationUrl      string    `json:"dbProcessUrl" db:"dbProcessUrl"`
	MerchantToken        string    `json:"merchantToken" db:"merchantToken"`
	RequestDate          string    `json:"reqDt" db:"reqDt"`
	RequestTime          string    `json:"reqTm" db:"reqTm"`
	RequestDomain        string    `json:"reqDomain" db:"reqDomain"`
	RequestServerIP      string    `json:"reqServerIP" db:"reqServerIP"`
	RequestClientVersion string    `json:"reqClientVer" db:"reqClientVer"`
	UserIP               string    `json:"userIP" db:"userIP"`
	UserSessionID        string    `json:"userSessionID" db:"userSessionID"`
	UserAgent            string    `json:"userAgent" db:"userAgent"`
	UserLanguage         string    `json:"userLanguage" db:"userLanguage"`
	CartData             string    `json:"cartData" db:"cartData"`
	InstallmentType      int       `json:"instmntType" db:"instmntType"`
	InstallmentMonth     int       `json:"instmntMon" db:"instmntMon"`
	RecurringOption      int       `json:"recurrOpt" db:"recurrOpt"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}
