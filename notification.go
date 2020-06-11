package nicepay

import "encoding/json"

type Notification struct {
	TxId            string          `json:"tXid" form:"tXid" query:"tXid"`
	MerchantToken   string          `json:"merchantToken" form:"merchantToken" query:"merchantToken"`
	ReferenceNo     string          `json:"referenceNo" form:"referenceNo" query:"referenceNo"`
	PaymentMethod   PaymentTypeCode `json:"payMethod" form:"payMethod" query:"payMethod"`
	Amount          json.Number     `json:"amt" form:"amt" query:"amt"`
	TransactionDate *NiceDate       `json:"transDt" form:"transDt" query:"transDt"`
	TransactionTime *NiceTime       `json:"transTm" form:"transTm" query:"transTm"`
	Currency        string          `json:"currency" form:"currency" query:"currency"`
	GoodsName       string          `json:"goodsNm" form:"goodsNm" query:"goodsNm"`
	BillingName     string          `json:"billingNm" form:"billingNm" query:"billingNm"`
	MatchFlag       int             `json:"matchCl" form:"matchCl" query:"matchCl"`
	Status          int             `json:"status" form:"status" query:"status"`

	*NotificationCreditCard
	*NotificationVa
	*NotificationOther
}

type NotificationCreditCard struct {
	AuthNo            string          `json:"authNo,omitempty" form:"authNo" query:"authNo"`
	IssueBankCode     BankCode        `json:"issuBankCd,omitempty" form:"issuBankCd" query:"issuBankCd"`
	IssueBankName     string          `json:"IssueBankNm,omitempty" form:"IssueBankNm" query:"IssueBankNm"`
	AcquireBankCode   BankCode        `json:"acquBankCd,omitempty" form:"acquBankCd" query:"acquBankCd"`
	AcquireBankName   string          `json:"acquBankNm,omitempty" form:"acquBankNm" query:"acquBankNm"`
	CardNo            string          `json:"cardNo,omitempty,omitempty" form:"cardNo" query:"cardNo"`
	CardExpiry        string          `json:"cardExpYymm,omitempty" form:"cardExpYymm" query:"cardExpYymm"`
	InstallmentMonth  string          `json:"instmntMon,omitempty" form:"instmntMon" query:"instmntMon"`
	InstallmentType   InstallmentType `json:"instmntType,omitempty" form:"instmntType" query:"instmntType"`
	PreAuthToken      string          `json:"preauthToken,omitempty" form:"preauthToken" query:"preauthToken"`
	RecurringToken    string          `json:"recurringToken,omitempty" form:"recurringToken" query:"recurringToken"`
	CcTransactionType string          `json:"ccTransType,omitempty" form:"ccTransType" query:"ccTransType"`
	Vat               json.Number     `json:"vat,omitempty" form:"vat" query:"vat"`
	Fee               json.Number     `json:"fee,omitempty" form:"fee" query:"fee"`
	NoTaxAmount       json.Number     `json:"notaxAmt,omitempty" form:"notaxAmt" query:"notaxAmt"`
}

type NotificationVa struct {
	BankCode     BankCode  `json:"bankCd" form:"bankCd" query:"bankCd"`
	VaNumber     string    `json:"vacctNo" form:"vacctNo" query:"vacctNo"`
	VAExpireDate *NiceDate `json:"vacctValidDt" form:"vacctValidDt" query:"vacctValidDt"`
	VAExpiryTime *NiceTime `json:"vacctValidTm" form:"vacctValidTm" query:"vacctValidTm"`
	DepositDate  *NiceDate `json:"depositDt" form:"depositDt" query:"depositDt"`
	DepositTime  *NiceTime `json:"depositTm" form:"depositTm" query:"depositTm"`
}

type NotificationOther struct {
	MitraCode    MitraCode `json:"mitraCd,omitempty" form:"mitraCd" query:"mitraCd"`
	PayNo        string    `json:"payNo,omitempty" form:"payNo" query:"payNo"`
	PayValidDate *NiceDate `json:"payValidDt,omitempty" form:"payValidDt" query:"payValidDt"`
	PayValidTime *NiceTime `json:"payValidTm,omitempty" form:"payValidTm" query:"payValidTm"`
	MRefNo       string    `json:"mRefNo,omitempty" form:"mRefNo" query:"mRefNo"`
	ReceiptCode  string    `json:"receiptCode,omitempty" form:"receiptCode" query:"receiptCode"`
	DepositDate  *NiceDate `json:"depositDt" form:"depositDt" query:"depositDt"`
	DepositTime  *NiceTime `json:"depositTm" form:"depositTm" query:"depositTm"`
}
