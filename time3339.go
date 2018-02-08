package src

import "time"

type Time3339 time.Time

const (
	salesforceFormat = time.RFC3339Nano
	customFormat     = "2006-01-02"
)

func defaultTime() Time3339 {
	return Time3339(time.Now())
}

func FromTime(t time.Time) Time3339 {
	return Time3339(t)
}

func NewFromYYYMMDD(input string) (Time3339, error) {
	t, err := time.Parse(customFormat, input)

	if err != nil {
		return defaultTime(), err
	}

	return Time3339(t), nil
}

func NewFromRFC3339(input string) (Time3339, error) {
	t, err := time.Parse(time.RFC3339, input)

	if err != nil {
		return defaultTime(), err
	}

	return Time3339(t), nil
}

func NewFromSalesforceFormat(input string) (Time3339, error) {
	t, err := time.Parse(salesforceFormat, input)

	if err != nil {
		return defaultTime(), err
	}

	return Time3339(t), nil
}

func (t *Time3339) ToRFC3339() string {
	return time.Time(*t).Format(time.RFC3339)
}

func (t *Time3339) ToSalesforceFormat() string {
	return time.Time(*t).Format(salesforceFormat)
}

func (t *Time3339) ToTime() time.Time {
	return time.Time(*t)
}

func (t *Time3339) ToUnix() int64 {
	return time.Time(*t).Unix()
}
