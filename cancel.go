package nicepay

type CancelRequest struct {
	Timestamp      NiceTimestamp   `json:"timeStamp"`
	TxId           string          `json:"tXid"`
	IMid           string          `json:"iMid"`
	PayMethod      PaymentTypeCode `json:"payMethod"`
	CancelType     CancelType      `json:"cancelType"`
	CancelMsg      string          `json:"cancelMsg,omitempty"`
	MerchantToken  string          `json:"merchantToken"`
	PreAuthToken   string          `json:"preauthToken"`
	Amount         float64         `json:"amt"`
	Fee            float64         `json:"fee,omitempty"`
	Vat            float64         `json:"vat,omitempty"`
	NoTaxAmount    float64         `json:"notaxAmt,omitempty"`
	CancelServerIp string          `json:"cancelServerIp,omitempty"`
	CancelUserId   string          `json:"cancelUserId,omitempty"`
	CancelUserIp   string          `json:"cancelUserIp,omitempty"`
	CancelUserInfo string          `json:"cancelUserInfo,omitempty"`
	CancelRetryCnt string          `json:"cancelRetryCnt,omitempty"`
	Worker         string          `json:"worker,omitempty"`
}

type CancelResponse struct {
	ResultCd            string    `json:"resultCd"`
	ResultMsg           string    `json:"resultMsg"`
	TxId                string    `json:"tXid,omitempty"`
	ReferenceNo         string    `json:"referenceNo,omitempty"`
	TransactionDate     *NiceDate `json:"transDt,omitempty"`
	TransactionTime     *NiceTime `json:"transTm,omitempty"`
	Desc                string    `json:"description,omitempty"`
	Amount              string    `json:"amt,omitempty"`
	CancelTransactionID string    `json:"canceltXid,omitempty"`
}
