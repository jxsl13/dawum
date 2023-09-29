package dto

import (
	"encoding/json"
	"net/url"
)

type link url.URL

func (p *link) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	u, err := url.ParseRequestURI(s)
	if err != nil {
		return err
	}

	*p = link(*u)
	return nil
}

func (p link) MarshalJSON() ([]byte, error) {
	u := url.URL(p)
	return json.Marshal(u.String())
}
