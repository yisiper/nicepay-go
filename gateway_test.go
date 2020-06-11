package nicepay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestGatewayClient(t *testing.T) {
	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		var responseString string
		switch req.URL.Path {
		case registrationPath, inquiryPath, paymentPath, cancelPath:
			responseString = `{"resultCd": "0000"}`
		default:
			responseString = `{"resultCd": "1004"}`
		}
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseString)),
			Header:     make(http.Header),
		}
	})
	client := NewClient("dummy", "dummy", "localhost")
	client.CustomHttpClient(httpClient)
	client.LogLevel = LogAll
	client.Logger.SetOutput(ioutil.Discard)

	gateway := NewCoreGateway(client)

	t.Run("registration should success", func(t *testing.T) {
		var err error
		payload := &RegistrationRequest{}

		_, err = gateway.Registration(payload)
		assert.NotNil(t, err)

		payload.Timestamp = NiceTimestamp{Time: time.Now()}
		_, err = gateway.Registration(payload)
		assert.NotNil(t, err)

		payload.ReferenceNo = "referenceNo"
		_, err = gateway.Registration(payload)
		assert.NotNil(t, err)

		payload.Amount = 100000
		req, err := gateway.Registration(payload)
		assert.NoError(t, err)
		assert.Equal(t, "0000", req.ResultCd)
	})

	t.Run("cancel should success", func(t *testing.T) {
		var err error
		payload := &CancelRequest{}

		_, err = gateway.Cancel(payload)
		assert.NotNil(t, err)

		payload.Timestamp = NiceTimestamp{Time: time.Now()}
		_, err = gateway.Cancel(payload)
		assert.NotNil(t, err)

		payload.TxId = "TxId"
		_, err = gateway.Cancel(payload)
		assert.NotNil(t, err)

		payload.Amount = 100000
		req, err := gateway.Cancel(payload)
		assert.NoError(t, err)
		assert.Equal(t, "0000", req.ResultCd)
	})

	t.Run("inquiry should success", func(t *testing.T) {
		var err error
		payload := &InquiryRequest{}
		_, err = gateway.Inquiry(payload)
		assert.NotNil(t, err)

		payload.Timestamp = NiceTimestamp{Time: time.Now()}
		_, err = gateway.Inquiry(payload)
		assert.NotNil(t, err)

		payload.ReferenceNo = "ReferenceNo"
		_, err = gateway.Inquiry(payload)
		assert.NotNil(t, err)

		payload.Amount = 100000
		req, err := gateway.Inquiry(payload)
		assert.NoError(t, err)
		assert.Equal(t, "0000", req.ResultCd)
	})

	t.Run("inquiry notification should success", func(t *testing.T) {
		req, err := gateway.InquiryFromNotification(&Notification{
			TxId:        "txId",
			ReferenceNo: "referenceNo",
			Amount:      json.Number("10000"),
		})
		assert.NoError(t, err)
		assert.Equal(t, "0000", req.ResultCd)
	})

	t.Run("payment should success", func(t *testing.T) {
		var err error
		payload := &PaymentRequest{}
		_, err = gateway.Payment(payload)
		assert.NotNil(t, err)

		payload.Timestamp = NiceTimestamp{Time: time.Now()}
		_, err = gateway.Payment(payload)
		assert.NotNil(t, err)

		payload.ReferenceNo = "referenceNo"
		_, err = gateway.Payment(payload)
		assert.NotNil(t, err)

		payload.Amount = 100000
		req, err := gateway.Payment(payload)
		assert.NoError(t, err)
		assert.Equal(t, "0000", req.ResultCd)
	})

	t.Run("iMid should error", func(t *testing.T) {
		var err error
		gateway.Client.IMid = ""
		_, err = gateway.Registration(&RegistrationRequest{
			Timestamp: NiceTimestamp{Time: time.Now()},
		})
		assert.NotNil(t, err)

		_, err = gateway.Inquiry(&InquiryRequest{
			Timestamp: NiceTimestamp{Time: time.Now()},
		})
		assert.NotNil(t, err)

		gateway.Client.IMid = client.IMid
	})

	t.Run("empty merchant key", func(t *testing.T) {
		gateway.Client.MerchantKey = ""
		_, err := gateway.Inquiry(&InquiryRequest{})
		assert.NotNil(t, err)
		gateway.Client.MerchantKey = client.MerchantKey
	})

	t.Run("fill merchant token has invalid type", func(t *testing.T) {
		err := gateway.fillMerchantToken("hello")
		assert.NotNil(t, err)
	})
}
