package opentrans

import (
	"errors"
	"time"

	"gitlab.com/mclgmbh/gomod/bmecat"
)

type multiLocaleString00050 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=50"`
}

const datetimeFormat = "2006-01-02T15:04:05+07:00"

type Datetime struct {
	time.Time
}

func (d Datetime) MarshalText() ([]byte, error) {
	if y := d.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(datetimeFormat))
	return d.Time.AppendFormat(b, datetimeFormat), nil
}

func (d *Datetime) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	d.Time, err = time.Parse(datetimeFormat, string(data))
	return err
}
