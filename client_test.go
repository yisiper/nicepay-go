package nicepay

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	server := httptest.NewServer(http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/error":
			time.Sleep(20 * time.Millisecond)
		case "/error-json":
			_, _ = w.Write([]byte(`"status": "ok"`))
		default:
			_, _ = w.Write([]byte(`{"status": "ok"}`))
		}
	}), 10*time.Millisecond, "server timeout"))

	defer server.Close()

	client := NewClient("dummy", "dummy", "localhost")
	client.CustomHttpClient(server.Client())
	client.LogLevel = LogAll
	client.Logger.SetOutput(ioutil.Discard)

	var err error
	var response map[string]interface{}

	err = client.Call("*?", server.URL, nil, &response)
	assert.NotNil(t, err)
	err = client.CallWithForm("*?", server.URL, nil, &response)
	assert.NotNil(t, err)

	err = client.Call("POST", server.URL, nil, &response)
	assert.NoError(t, err)

	err = client.CallWithForm("POST", server.URL, nil, &response)
	assert.NoError(t, err)

	err = client.Call("POST", server.URL+"/error", nil, &response)
	assert.NotNil(t, err)
}
