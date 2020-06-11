package nicepay

import "encoding/json"

type InquiryRequest struct {
	Timestamp     NiceTimestamp `json:"timeStamp"`
	TxId          string        `json:"tXid"`
	IMid          string        `json:"iMid"`
	ReferenceNo   string        `json:"referenceNo"`
	Amount        float64       `json:"amt"`
	MerchantToken string        `json:"merchantToken"`
}

type InquiryResponse struct {
	ResultCd        string          `json:"resultCd"`
	ResultMsg       string          `json:"resultMsg"`
	TxId            string          `json:"tXid"`
	IMid            string          `json:"iMid"`
	ReferenceNo     string          `json:"referenceNo"`
	PayMethod       PaymentTypeCode `json:"payMethod"`
	Amount          json.Number     `json:"amt"`
	CancelAmt       string          `json:"cancelAmt"`
	ReqDt           *NiceDate       `json:"reqDt"`
	ReqTm           *NiceTime       `json:"reqTm"`
	TransactionDate *NiceDate       `json:"transDt"`
	TransactionTime *NiceTime       `json:"transTm"`
	DepositDate     *NiceDate       `json:"depositDt"`
	DepositTime     *NiceTime       `json:"depositTm"`
	Currency        string          `json:"currency"`
	GoodsName       string          `json:"goodsNm"`
	BillingName     string          `json:"billingNm"`
	Status          string          `json:"status"`

	*InquiryVaResponse
	*InquiryCreditCardResponse
	*InquiryOtherResponse
}

type InquiryCreditCardResponse struct {
	AuthNo            string          `json:"authNo,omitempty"`
	IssueBankCode     BankCode        `json:"issuBankCd,omitempty"`
	AcquireBankCode   BankCode        `json:"acquBankCd,omitempty"`
	CardNo            string          `json:"cardNo,omitempty"`
	CardExpiry        string          `json:"cardExpYymm,omitempty"`
	InstallmentMonth  string          `json:"instmntMon,omitempty"`
	InstallmentType   InstallmentType `json:"instmntType,omitempty"`
	PreAuthToken      string          `json:"preauthToken,omitempty"`
	RecurringToken    string          `json:"recurringToken,omitempty"`
	CcTransactionType string          `json:"ccTransType,omitempty"`
	AcquireStatus     string          `json:"acquStatus,omitempty"`
	Vat               json.Number     `json:"vat,omitempty"`
	Fee               json.Number     `json:"fee,omitempty"`
	NoTaxAmount       json.Number     `json:"notaxAmt,omitempty"`
}

type InquiryVaResponse struct {
	BankCode     BankCode  `json:"bankCd,omitempty"`
	VaNumber     string    `json:"vacctNo,omitempty"`
	VAExpireDate *NiceDate `json:"vacctValidDt,omitempty"`
	VAExpiryTime *NiceTime `json:"vacctValidTm,omitempty"`
}

type InquiryOtherResponse struct {
	MitraCode    MitraCode `json:"mitraCd,omitempty"`
	PayNo        string    `json:"payNo,omitempty"`
	PayValidDate *NiceDate `json:"payValidDt,omitempty"`
	PayValidTime *NiceTime `json:"payValidTm,omitempty"`
	MRefNo       string    `json:"mRefNo,omitempty"`
	ReceiptCode  string    `json:"receiptCode,omitempty"`
}
