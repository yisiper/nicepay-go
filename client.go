package nicepay

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/pkg/errors"
)

const httpTimeout = 30 * time.Second

var httpClient = &http.Client{Timeout: httpTimeout}

type Client struct {
	IMid        string
	MerchantKey string
	CallbackUrl string
	Env         Environment
	httpClient  *http.Client
	LogLevel    LogLevel
	Logger      *log.Logger
}

type LogLevel int

const (
	LogNothing LogLevel = iota
	LogInfo
	LogError
	LogAll
)

func NewClient(iMid string, merchantKey string, callback string) *Client {
	return &Client{
		IMid:        iMid,
		MerchantKey: merchantKey,
		Env:         Development,
		CallbackUrl: callback,
		LogLevel:    LogAll,
		Logger:      log.New(os.Stderr, "", log.LstdFlags),
		httpClient:  httpClient,
	}
}

func (c *Client) CustomHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.log(LogError, err)
		return nil, errors.Wrap(err, "new http request context failed")
	}
	return req, nil
}

func (c *Client) Call(method, path string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		c.log(LogError, err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return c.ExecuteRequest(req, v)
}

func (c *Client) CallWithForm(method, path string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		c.log(LogError, err)
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.ExecuteRequest(req, v)
}

func (c *Client) ExecuteRequest(r *http.Request, v interface{}) error {
	c.log(LogInfo, fmt.Sprintf("Request %s: %s %s", r.Method, r.URL.Host, r.URL.Path))

	res, err := c.httpClient.Do(r)
	if err != nil {
		c.log(LogError, err)
		return errors.Wrap(err, "http request failed")
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return errors.Wrap(err, reflect.TypeOf(v).String())
	}
	return nil
}

func (c *Client) log(l LogLevel, msg interface{}) {
	if c.LogLevel >= l && msg != nil {
		c.Logger.Println(msg)
	}
}
