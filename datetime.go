package nicepay

import (
	"fmt"
	"strings"
	"time"
)

const (
	NiceDateFormat      = "20060102"       // YYYYMMDD
	NiceTimeFormat      = "150405"         // HH24MISS
	NiceTimestampFormat = "20060102150405" //20180123100505
)

type NiceDate struct {
	time.Time
}

type NiceTime struct {
	time.Time
}

type NiceTimestamp struct {
	time.Time
}

func (t NiceDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(NiceDateFormat))
	return []byte(stamp), nil
}

func (t *NiceDate) UnmarshalJSON(b []byte) error {
	var err error
	stringJson := strings.Trim(string(b), "\"")
	t.Time, err = time.Parse(NiceDateFormat, stringJson)
	return err
}

func (t *NiceDate) UnmarshalParam(src string) error {
	ts, err := time.Parse(NiceDateFormat, src)
	*t = NiceDate{Time: ts}
	return err
}

func (t *NiceDate) String() string {
	return t.Format(NiceDateFormat)
}

func (t NiceTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(NiceTimeFormat))
	return []byte(stamp), nil
}

func (t *NiceTime) UnmarshalJSON(b []byte) error {
	var err error
	stringJson := strings.Trim(string(b), "\"")
	t.Time, err = time.Parse(NiceTimeFormat, stringJson)
	return err
}

func (t *NiceTime) UnmarshalParam(src string) error {
	ts, err := time.Parse(NiceTimeFormat, src)
	*t = NiceTime{Time: ts}
	return err
}

func (t *NiceTime) String() string {
	return t.Format(NiceTimeFormat)
}

func (t NiceTimestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(NiceTimestampFormat))
	return []byte(stamp), nil
}

func (t *NiceTimestamp) UnmarshalJSON(b []byte) error {
	var err error
	stringJson := strings.Trim(string(b), "\"")
	t.Time, err = time.Parse(NiceTimestampFormat, stringJson)
	return err
}

func (t *NiceTimestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(NiceTimestampFormat, src)
	*t = NiceTimestamp{Time: ts}
	return err
}

func (t *NiceTimestamp) String() string {
	return t.Format(NiceTimestampFormat)
}
