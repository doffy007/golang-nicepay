package constants

type PaymentMethod string

const (
	MethodCreditCard       PaymentMethod = "01"
	MethodVirtualAccount   PaymentMethod = "02"
	MethodConvenienceStore PaymentMethod = "03"
	MethodClickPay         PaymentMethod = "04"
	MethodEWallet          PaymentMethod = "05"
)

var AllPaymentMethod = []PaymentMethod{
	MethodCreditCard,
	MethodVirtualAccount,
	MethodConvenienceStore,
	MethodClickPay,
	MethodEWallet,
}

type InstallmentType int

const (
	CustomerCharge InstallmentType = 1
	MerchantCharge InstallmentType = 2
)
