package dawum

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jxsl13/dawum/dto"
)

var (
	apiUrl = "https://api.dawum.de/"
)

func GetData(ctx context.Context, os ...Option) (*dto.Data, error) {
	opts := options{
		client: http.DefaultClient,
	}

	for _, o := range os {
		err := o(&opts)
		if err != nil {
			return nil, fmt.Errorf("invalid option: %w", err)
		}
	}

	return getDataWithClient(ctx, http.DefaultClient)
}

func getDataWithClient(ctx context.Context, client *http.Client) (*dto.Data, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "https://github.com/jxsl13/dawum")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("failed to get api data: %s", resp.Status)
	}

	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(reader)
		if err != nil {
			return nil, fmt.Errorf("invalid response encoding: expected gzip: %w", err)
		}
	}

	var data dto.Data
	err = json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode api data: %w", err)
	}

	return &data, nil
}

func GetLastUpdate(ctx context.Context, os ...Option) (time.Time, error) {
	opts := options{
		client: http.DefaultClient,
	}

	for _, o := range os {
		err := o(&opts)
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid option: %w", err)
		}
	}

	return getLastUpdateWithClient(ctx, opts.client)
}

func getLastUpdateWithClient(ctx context.Context, client *http.Client) (time.Time, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl+"/last_update.txt", nil)
	if err != nil {
		return time.Time{}, err
	}

	req.Header.Set("Accept", "text/plain; charset=utf-8")
	req.Header.Set("User-Agent", "https://github.com/jxsl13/dawum")

	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return time.Time{}, fmt.Errorf("failed to get api last update: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to read api last update: %w", err)
	}

	t, err := time.Parse(dto.W3CDateTimeLayout, string(data))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse api last update: %w", err)
	}

	return t, nil
}
