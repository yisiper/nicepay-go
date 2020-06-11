package nicepay

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const defaultFormat = "2006-01-02 15:04:05"

type dateTimeTest struct {
	From       string
	To         string
	FromFormat string
	ToFormat   string
	Json       string
}

type dateTimeInterface interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
	UnmarshalParam(string) error
	String() string
}

var dateTimeTests = []dateTimeTest{
	{
		From:       "2020-01-01 12:12:12",
		To:         "20200101",
		FromFormat: defaultFormat,
		ToFormat:   NiceDateFormat,
		Json:       `"20200101"`,
	},
	{
		From:       "2020-01-01 12:12:12",
		To:         "121212",
		FromFormat: defaultFormat,
		ToFormat:   NiceTimeFormat,
		Json:       `"121212"`,
	}, {
		From:       "2020-01-01 12:12:12",
		To:         "20200101121212",
		FromFormat: defaultFormat,
		ToFormat:   NiceTimestampFormat,
		Json:       `"20200101121212"`,
	},
}

func TestMain(m *testing.M) {
	niceDateTimeValidate()
	os.Exit(m.Run())
}

func TestNiceDate(t *testing.T) {
	for _, test := range dateTimeTests {
		if test.FromFormat != NiceDateFormat && test.ToFormat != NiceDateFormat {
			continue
		}
		from, _ := time.Parse(test.FromFormat, test.From)
		d := &NiceDate{Time: from}
		niceDateTimeRunner(t, test, d)
	}
}

func TestNiceTime(t *testing.T) {
	for _, test := range dateTimeTests {
		if test.FromFormat != NiceTimeFormat && test.ToFormat != NiceTimeFormat {
			continue
		}
		from, _ := time.Parse(test.FromFormat, test.From)
		d := &NiceTime{Time: from}
		niceDateTimeRunner(t, test, d)
	}
}

func TestNiceTimestamp(t *testing.T) {
	for _, test := range dateTimeTests {
		if test.FromFormat != NiceTimestampFormat && test.ToFormat != NiceTimestampFormat {
			continue
		}
		from, _ := time.Parse(test.FromFormat, test.From)
		d := &NiceTimestamp{Time: from}
		niceDateTimeRunner(t, test, d)
	}
}

func niceDateTimeValidate() {
	for _, test := range dateTimeTests {
		if _, err := time.Parse(test.FromFormat, test.From); err != nil {
			log.Fatal(err)
		}
		if _, err := time.Parse(test.ToFormat, test.To); err != nil {
			log.Fatal(err)
		}
	}
}

func niceDateTimeRunner(t *testing.T, test dateTimeTest, d dateTimeInterface) {
	t.Run("stringer", func(t *testing.T) {
		assert.EqualValues(t, test.To, d.String())
	})

	t.Run("marshal json", func(t *testing.T) {
		got, err := d.MarshalJSON()
		assert.NoError(t, err)
		assert.EqualValues(t, test.Json, got)
	})

	t.Run("unmarshal json", func(t *testing.T) {
		err := d.UnmarshalJSON([]byte(test.Json))
		assert.NoError(t, err)
		assert.EqualValues(t, test.To, d.String())
	})

	t.Run("unmarshal param", func(t *testing.T) {
		err := d.UnmarshalParam(test.To)
		assert.NoError(t, err)
		assert.EqualValues(t, test.To, d.String())
	})
}
