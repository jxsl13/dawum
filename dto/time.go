package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	W3CDateTimeLayout = "2006-01-02T15:04:05-07:00"
	DateOnlyLayout    = "2006-01-02"
)

// 2023-09-29T09:41:59+02:00
type dateTime time.Time

func (p *dateTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(W3CDateTimeLayout, s)
	if err != nil {
		return fmt.Errorf("invalid time format: expected: %s: got: %s", W3CDateTimeLayout, s)
	}

	*p = dateTime(t)
	return nil
}

func (p dateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(p).Format(W3CDateTimeLayout))
}

type date time.Time

func (p *date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(DateOnlyLayout, s)
	if err != nil {
		return fmt.Errorf("invalid time format: expected: %s: got: %s", DateOnlyLayout, s)
	}

	*p = date(t)
	return nil
}

func (p date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(p).Format(DateOnlyLayout))
}
