package nicepay

type InstallmentType string

const (
	InstallmentCustomerCharge InstallmentType = "1"
	InstallmentMerchantCharge                 = "2"
)

type PaymentTypeCode string

const (
	SourceCreditCard       PaymentTypeCode = "01"
	SourceVirtualAccount                   = "02"
	SourceConvenienceStore                 = "03"
	// SourceClickPay                           = "04"
	// SourceEWallet                            = "05"
)

type PaymentStatusCode int

const (
	CreditSuccess  PaymentStatusCode = 0
	CreditFailed                     = 1
	CreditVoid                       = 2
	CreditReversal                   = 3
	VaPaid                           = 0
	VaUnpaid                         = 3
	VaExpired                        = 4
	CvsPaid                          = 0
	CvsUnpaid                        = 3
	CvsExpired                       = 4
	CvsReadyToPaid                   = 5
)

type NotificationStatusCode int

const (
	NotificationMatch NotificationStatusCode = 1
	NotificationOver                         = 2
	NotificationUnder                        = 3
)

type CancelType int

const (
	FullCancellation    CancelType = 1
	PartialCancellation            = 2
)

type CreditCardTransactionType int

const (
	CreditCardTransactionNormal    CreditCardTransactionType = 1
	CreditCardTransactionRecurring                           = 2
	CreditCardTransactionPreAuth                             = 3
	CreditCardTransactionCaptured                            = 4
)
