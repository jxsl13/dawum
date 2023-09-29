package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	data := []byte(`"2023-09-29T09:41:59+02:00"`)

	var p dateTime
	err := json.Unmarshal(data, &p)
	require.NoError(t, err)

	b, err := json.Marshal(p)
	require.NoError(t, err)
	require.Equal(t, data, b)
}

func TestDate(t *testing.T) {
	data := []byte(`"2023-09-29"`)

	var p date
	err := json.Unmarshal(data, &p)
	require.NoError(t, err)

	b, err := json.Marshal(p)
	require.NoError(t, err)
	require.Equal(t, data, b)
}
