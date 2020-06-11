package nicepay

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/pkg/errors"
)

const (
	registrationPath = "/nicepay/direct/v2/registration"
	inquiryPath      = "/nicepay/direct/v2/inquiry"
	paymentPath      = "/nicepay/direct/v2/payment"
	cancelPath       = "/nicepay/direct/v2/cancel"
)

type Gateway struct {
	Client *Client
}

func NewCoreGateway(c *Client) *Gateway {
	return &Gateway{
		Client: c,
	}
}

func (g *Gateway) Call(method string, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return g.Client.Call(method, g.Client.Env.String()+path, body, v)
}

func (g *Gateway) CallWithForm(method string, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return g.Client.CallWithForm(method, g.Client.Env.String()+path, body, v)
}

func (g *Gateway) Registration(req *RegistrationRequest) (*RegistrationResponse, error) {
	resp := &RegistrationResponse{}
	if err := g.fillMerchantToken(req); err != nil {
		return nil, err
	}

	if req.DbProcessUrl == "" {
		req.DbProcessUrl = g.Client.CallbackUrl
	}

	b, _ := json.Marshal(req)
	if err := g.Call(http.MethodPost, registrationPath, bytes.NewReader(b), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Gateway) Cancel(req *CancelRequest) (*CancelResponse, error) {
	resp := &CancelResponse{}
	if err := g.fillMerchantToken(req); err != nil {
		return nil, err
	}

	b, _ := json.Marshal(req)
	if err := g.Call(http.MethodPost, cancelPath, bytes.NewReader(b), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Gateway) Payment(req *PaymentRequest) (*PaymentResponse, error) {
	resp := &PaymentResponse{}
	if err := g.fillMerchantToken(req); err != nil {
		return nil, err
	}

	v, _ := query.Values(req)
	if err := g.CallWithForm(http.MethodPost, paymentPath, bytes.NewBufferString(v.Encode()), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Gateway) Inquiry(req *InquiryRequest) (*InquiryResponse, error) {
	resp := &InquiryResponse{}
	if err := g.fillMerchantToken(req); err != nil {
		return nil, err
	}

	b, _ := json.Marshal(req)
	if err := g.Call(http.MethodPost, inquiryPath, bytes.NewReader(b), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Gateway) InquiryFromNotification(notification *Notification) (*InquiryResponse, error) {
	amount, _ := notification.Amount.Float64()
	return g.Inquiry(&InquiryRequest{
		TxId:        notification.TxId,
		Amount:      amount,
		ReferenceNo: notification.ReferenceNo,
		Timestamp:   NiceTimestamp{Time: time.Now()},
		IMid:        g.Client.IMid,
	})
}

func (g *Gateway) fillMerchantToken(v interface{}) error {
	var err error
	if g.Client.MerchantKey == "" {
		return errors.New("merchant key is empty")
	}

	sha := sha256.New()
	switch vType := v.(type) {
	case *RegistrationRequest:
		vType.IMid = g.Client.IMid
		err = errors.New("generate token registration")
		if vType.Timestamp.IsZero() {
			return errors.Wrap(err, "invalid timestamp")
		}
		if vType.IMid == "" {
			return errors.Wrap(err, "invalid iMid")
		}
		if vType.ReferenceNo == "" {
			return errors.Wrap(err, "invalid reference code")
		}
		if vType.Amount == 0 {
			return errors.Wrap(err, "invalid amount")
		}
		sha.Write([]byte(vType.Timestamp.String()))
		sha.Write([]byte(vType.IMid))
		sha.Write([]byte(vType.ReferenceNo))
		sha.Write([]byte(floatToStr(vType.Amount)))
		sha.Write([]byte(g.Client.MerchantKey))
		vType.MerchantToken = hex.EncodeToString(sha.Sum(nil))
		return nil

	case *InquiryRequest:
		vType.IMid = g.Client.IMid
		err = errors.New("generate token inquiry")
		if vType.Timestamp.IsZero() {
			return errors.Wrap(err, "invalid timestamp")
		}
		if vType.IMid == "" {
			return errors.Wrap(err, "invalid iMid")
		}
		if vType.ReferenceNo == "" {
			return errors.Wrap(err, "invalid reference code")
		}
		if vType.Amount == 0 {
			return errors.Wrap(err, "invalid amount")
		}
		sha.Write([]byte(vType.Timestamp.String()))
		sha.Write([]byte(vType.IMid))
		sha.Write([]byte(vType.ReferenceNo))
		sha.Write([]byte(floatToStr(vType.Amount)))
		sha.Write([]byte(g.Client.MerchantKey))
		vType.MerchantToken = hex.EncodeToString(sha.Sum(nil))
		return nil

	case *PaymentRequest:
		vType.IMid = g.Client.IMid
		err = errors.New("generate payment")
		if vType.Timestamp.IsZero() {
			return errors.Wrap(err, "invalid timestamp")
		}
		if vType.ReferenceNo == "" {
			return errors.Wrap(err, "invalid reference code")
		}
		if vType.Amount == 0 {
			return errors.Wrap(err, "invalid amount")
		}
		sha.Write([]byte(vType.Timestamp.String()))
		sha.Write([]byte(vType.IMid))
		sha.Write([]byte(vType.ReferenceNo))
		sha.Write([]byte(floatToStr(vType.Amount)))
		sha.Write([]byte(g.Client.MerchantKey))
		vType.MerchantToken = hex.EncodeToString(sha.Sum(nil))
		return nil

	case *CancelRequest:
		vType.IMid = g.Client.IMid
		err = errors.New("generate cancel request")
		if vType.Timestamp.IsZero() {
			return errors.Wrap(err, "invalid timestamp")
		}
		if vType.TxId == "" {
			return errors.Wrap(err, "invalid transaction id")
		}
		if vType.Amount == 0 {
			return errors.Wrap(err, "invalid amount")
		}
		sha.Write([]byte(vType.Timestamp.String()))
		sha.Write([]byte(vType.IMid))
		sha.Write([]byte(vType.TxId))
		sha.Write([]byte(floatToStr(vType.Amount)))
		sha.Write([]byte(g.Client.MerchantKey))
		vType.MerchantToken = hex.EncodeToString(sha.Sum(nil))
		return nil

	default:
		err = errors.New("no type found")
	}

	return err
}

func floatToStr(num float64) string {
	s := fmt.Sprintf("%.6f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}
