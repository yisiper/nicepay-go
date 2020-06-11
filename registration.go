package nicepay

import "encoding/json"

type RegistrationRequest struct {
	Timestamp      NiceTimestamp   `json:"timeStamp"`
	IMid           string          `json:"iMid"`
	PayMethod      PaymentTypeCode `json:"payMethod"`
	Currency       string          `json:"currency"`
	Amount         float64         `json:"amt"`
	ReferenceNo    string          `json:"referenceNo"`
	GoodsName      string          `json:"goodsNm"`
	BillingName    string          `json:"billingNm"`
	BillingPhone   string          `json:"billingPhone"`
	BillingEmail   string          `json:"billingEmail"`
	BillingAddr    string          `json:"billingAddr,omitempty"`
	BillingCity    string          `json:"billingCity"`
	BillingState   string          `json:"billingState"`
	BillingPostCd  string          `json:"billingPostCd"`
	BillingCountry string          `json:"billingCountry"`
	DbProcessUrl   string          `json:"dbProcessUrl"`
	CartData       string          `json:"cartData"`
	UserIp         string          `json:"userIP"`
	Description    string          `json:"description"`
	MerchantToken  string          `json:"merchantToken"`

	*RegistrationCreditCardRequest
	*RegistrationVaRequest
	*RegistrationCVSRequest

	DeliveryName     *string   `json:"deliveryNm,omitempty"`
	DeliveryPhone    *string   `json:"deliveryPhone,omitempty"`
	DeliveryAddr     *string   `json:"deliveryAddr,omitempty"`
	DeliveryCity     *string   `json:"deliveryCity,omitempty"`
	DeliveryState    *string   `json:"deliveryState,omitempty"`
	DeliveryPostCd   *string   `json:"deliveryPostCd,omitempty"`
	DeliveryCountry  *string   `json:"deliveryCountry,omitempty"`
	Vat              *float64  `json:"vat,omitempty"`
	Fee              *float64  `json:"fee,omitempty"`
	NoTaxAmount      *float64  `json:"notaxAmt,omitempty"`
	ReqDate          *NiceDate `json:"reqDt,omitempty"`
	ReqTime          *NiceTime `json:"reqTm,omitempty"`
	ReqDomain        *string   `json:"reqDomain,omitempty"`
	ReqServerIP      *string   `json:"reqServerIP,omitempty"`
	ReqClientVersion *string   `json:"reqClientVer,omitempty"`
	UserSessionID    *string   `json:"userSessionID,omitempty"`
	UserAgent        *string   `json:"userAgent,omitempty"`
	UserLanguage     *string   `json:"userLanguage,omitempty"`
}

type RegistrationCreditCardRequest struct {
	InstallmentType  InstallmentType `json:"instmntType"`
	InstallmentMonth int             `json:"instmntMon"`
	RecurringOption  *int            `json:"recurrOpt"`
}

type RegistrationVaRequest struct {
	BankCode             BankCode `json:"bankCd"`
	VAExpireDate         NiceDate `json:"vacctValidDt"`
	VAExpiryTime         NiceTime `json:"vacctValidTm"`
	MerchantReservedVaId *string  `json:"merFixAcctId,omitempty"`
}

type RegistrationCVSRequest struct {
	MitraCode MitraCode `json:"mitraCd"`
}

type CartData struct {
	Count int `json:"count"`
	Item  struct {
		ImageUrl string  `json:"img_url"`
		Name     string  `json:"goods_name"`
		Detail   string  `json:"goods_detail"`
		Amount   float64 `json:"goods_amt"`
	} `json:"struct"`
}

type RegistrationResponse struct {
	ResultCd        string          `json:"resultCd"`
	ResultMsg       string          `json:"resu	ltMsg"`
	TXid            string          `json:"tXid"`
	ReferenceNo     string          `json:"referenceNo"`
	PayMethod       PaymentTypeCode `json:"payMethod"`
	Amount          json.Number     `json:"amt"`
	TransactionDate *NiceDate       `json:"transDt"`
	TransactionTime *NiceTime       `json:"transTm"`
	Description     string          `json:"description"`
	Currency        string          `json:"currency"`
	GoodsName       string          `json:"goodsNm"`
	BillingName     string          `json:"billingNm"`
	BankCode        BankCode        `json:"bankCd"`
	VaNumber        string          `json:"vacctNo"`
	VAExpireDate    *NiceDate       `json:"vacctValidDt"`
	VAExpiryTime    *NiceTime       `json:"vacctValidTm"`
	MitraCode       MitraCode       `json:"mitraCd"`
	PayNo           string          `json:"payNo"`
	PayValidDt      *NiceDate       `json:"payValidDt"`
	PayValidTm      *NiceTime       `json:"payValidTm"`
	RequestURL      string          `json:"requestURL"`
}
