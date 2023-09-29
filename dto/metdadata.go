package dto

import (
	"encoding/json"
	"net/url"
	"time"
)

type Database struct {
	License    License   `json:"License"`
	Publisher  string    `json:"Publisher"`
	Author     string    `json:"Author"`
	LastUpdate time.Time `json:"Last_Update"`
}

func (d *Database) UnmarshalJSON(b []byte) error {
	var primary struct {
		License    License  `json:"License"`
		Publisher  string   `json:"Publisher"`
		Author     string   `json:"Author"`
		LastUpdate dateTime `json:"Last_Update"`
	}
	err := json.Unmarshal(b, &primary)
	if err != nil {
		return err
	}

	*d = Database{
		License:    primary.License,
		Publisher:  primary.Publisher,
		Author:     primary.Author,
		LastUpdate: time.Time(primary.LastUpdate),
	}
	return nil
}

func (d Database) MarshalJSON() ([]byte, error) {
	primary := struct {
		License    License  `json:"License"`
		Publisher  string   `json:"Publisher"`
		Author     string   `json:"Author"`
		LastUpdate dateTime `json:"Last_Update"`
	}{
		License:    d.License,
		Publisher:  d.Publisher,
		Author:     d.Author,
		LastUpdate: dateTime(d.LastUpdate),
	}
	return json.Marshal(primary)
}

type License struct {
	Name     string  `json:"Name"`
	Shortcut string  `json:"Shortcut"`
	Link     url.URL `json:"Link"`
}

func (l *License) UnmarshalJSON(b []byte) error {
	var primary struct {
		Name     string `json:"Name"`
		Shortcut string `json:"Shortcut"`
		Link     link   `json:"Link"`
	}
	err := json.Unmarshal(b, &primary)
	if err != nil {
		return err
	}

	*l = License{
		Name:     primary.Name,
		Shortcut: primary.Shortcut,
		Link:     url.URL(primary.Link),
	}
	return nil
}

func (l License) MarshalJSON() ([]byte, error) {
	primary := struct {
		Name     string `json:"Name"`
		Shortcut string `json:"Shortcut"`
		Link     link   `json:"Link"`
	}{
		Name:     l.Name,
		Shortcut: l.Shortcut,
		Link:     link(l.Link),
	}
	return json.Marshal(primary)
}
