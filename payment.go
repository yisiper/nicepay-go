package nicepay

type PaymentRequest struct {
	Timestamp      NiceTimestamp `json:"timeStamp"`
	TxId           string        `json:"tXid"`
	CardNo         string        `json:"cardNo,omiempty"`
	CardExpiry     string        `json:"cardExpYymm,omiempty"`
	CardCvv        string        `json:"cardCvv,omiempty"`
	CardHolderName string        `json:"cardHolderNm,omiempty"`
	RecurringToken string        `json:"recurringToken,omiempty"`
	PreAuthToken   string        `json:"preauthToken,omiempty"`
	ClickPayNo     string        `json:"clickPayNo,omiempty"`
	DataField3     string        `json:"dataField3,omiempty"`
	ClickPayToken  string        `json:"clickPayToken,omiempty"`
	MerchantToken  string        `json:"merchantToken"`
	CallBackUrl    string        `json:"callBackUrl"`
	IMid           string        `json:"-" url:"-"`
	ReferenceNo    string        `json:"-" url:"-"`
	Amount         float64       `json:"-" url:"-"`
}

type PaymentResponse struct {
	ResultCd          string          `json:"resultCd"`
	ResultMsg         string          `json:"resultMsg"`
	TxId              string          `json:"tXid"`
	ReferenceNo       string          `json:"referenceNo"`
	PaymentMethod     PaymentTypeCode `json:"payMethod"`
	Amount            float64         `json:"amt"`
	Currency          string          `json:"currency"`
	GoodsName         string          `json:"goodsNm"`
	BillingName       string          `json:"billingNm"`
	TransactionDate   *NiceDate       `json:"transDt"`
	TransactionTime   *NiceTime       `json:"transTm"`
	Description       string          `json:"description"`
	AuthNo            string          `json:"authNo"`
	IssueBankCode     BankCode        `json:"issuBankCd"`
	IssueBankName     string          `json:"issuBankNm"`
	AcquiringBankCode BankCode        `json:"acquBankCd"`
	AcquiringBankName string          `json:"acquBankNm"`
	CardNo            string          `json:"cardNo"`
	CardExpiry        string          `json:"cardExpYymm"`
	InstallmentMonth  int             `json:"instmntMon"`
	InstallmentType   InstallmentType `json:"instmntType"`
	RecurringToken    string          `json:"recurringToken"`
	PreAuthToken      string          `json:"preauthToken"`
	CCTransactionType string          `json:"ccTransType"`
	Vat               float64         `json:"vat"`
	Fee               float64         `json:"fee"`
	NoTaxAmount       float64         `json:"notaxAmt"`
	MitraCode         string          `json:"mitraCd"`
	ReceiptCode       string          `json:"receiptCode"`
	MRefNo            string          `json:"mRefNo"`
}
