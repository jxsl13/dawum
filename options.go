package dawum

import "net/http"

func WithClient(client *http.Client) Option {
	return func(o *options) error {
		o.client = client
		return nil
	}
}

type Option func(*options) error

type options struct {
	client *http.Client
}
