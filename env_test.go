package nicepay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment_String(t *testing.T) {
	data := []struct {
		Name string
		Env  Environment
		Url  string
	}{
		{
			Name: "production",
			Env:  Production,
			Url:  "https://api.nicepay.co.id",
		},
		{
			Name: "development",
			Env:  Development,
			Url:  "https://dev.nicepay.co.id",
		},
		{
			Name: "none",
			Url:  "",
		},
	}

	for _, d := range data {
		t.Run(d.Name, func(t *testing.T) {
			assert.Equal(t, d.Env.String(), d.Url, "environment url not equal")
		})
	}
}
