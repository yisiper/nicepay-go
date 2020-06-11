package nicepay

import (
	"encoding/json"
	"testing"
)

func TestResponse(t *testing.T) {
	j := `{"resultCd":"0000","resultMsg":"SUCCESS","tXid":"BMRITEST0101202005201128489898","referenceNo":"ORD12345","payMethod":"01","amt":"10000","transDt":"20200520","transTm":"112848","description":"","bankCd":null,"vacctNo":null,"mitraCd":null,"payNo":null,"currency":null,"goodsNm":null,"billingNm":null,"vacctValidDt":null,"vacctValidTm":null,"payValidDt":null,"payValidTm":null,"requestURL":null}`

	var resp RegistrationResponse
	if err := json.Unmarshal([]byte(j), &resp); err != nil {
		t.Error(err)
	}
}

func TestCancelResponse(t *testing.T) {
	j := `{"tXid":"TESTIDTEST01201803051530400331","referenceNo":"ORD12345","resultCd":"0000","resultMsg":"SUCCESS","transDt":"20180305","transTm":"153040","description":"Order Description","amt":"10000"}`
	var resp CancelResponse
	if err := json.Unmarshal([]byte(j), &resp); err != nil {
		t.Error(err)
	}
}

func TestInquiryResponse(t *testing.T) {
	j := `{"tXid":"TESTIDTEST05201803051150375209","iMid":"TESTIDTEST","currency":"IDR","amt":"10000","instmntMon":null,"instmntType":"1","referenceNo":"ORD12345","goodsNm":"Test Transaction Nicepay","payMethod":"05","billingNm":"Customer Name","reqDt":"20180305","reqTm":"115037","status":"9","resultCd":"0000","resultMsg":"init","cardNo":null,"preauthToken":null,"acquBankCd":null,"issuBankCd":null,"vacctValidDt":null,"vacctValidTm":null,"vacctNo":null,"bankCd":null,"payNo":null,"mitraCd":null,"receiptCode":null,"cancelAmt":null,"transDt":null,"transTm":null,"recurringToken":null,"ccTransType":null,"payValidDt":null,"payValidTm":null,"mRefNo":null,"acquStatus":null,"cardExpYymm":null,"acquBankNm":null,"issuBankNm":null,"depositDt":null,"depositTm":null}`

	var resp InquiryResponse
	if err := json.Unmarshal([]byte(j), &resp); err != nil {
		t.Fatal(err)
	}
}
