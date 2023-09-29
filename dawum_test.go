package dawum_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jxsl13/dawum"
	"github.com/jxsl13/dawum/dto"
	"github.com/stretchr/testify/require"
)

func TestHTTPGet(t *testing.T) {
	latestAtTest, err := time.Parse(dto.W3CDateTimeLayout, "2023-09-29T00:00:00+02:00")
	require.NoError(t, err)

	d, err := dawum.GetData(context.TODO())
	require.NoError(t, err)

	require.True(t, d.Database.LastUpdate.After(latestAtTest))
	require.NotEmpty(t, d.Parties)
	require.NotEmpty(t, d.Institutes)
	require.NotEmpty(t, d.Parliaments)
	require.NotEmpty(t, d.Surveys)
	require.NotEmpty(t, d.Taskers)
	require.NotEmpty(t, d.Methods)
}

func TestLastUpdate(t *testing.T) {
	latestAtTest, err := time.Parse(dto.W3CDateTimeLayout, "2023-09-29T00:00:00+02:00")
	require.NoError(t, err)

	lastUpdate, err := dawum.GetLastUpdate(context.Background())
	require.NoError(t, err)

	require.True(t, lastUpdate.After(latestAtTest))
}

func ExampleGetData() {
	data, err := dawum.GetData(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	// {"Database":{"License":{"Name":"ODC Open Database License","Shortcut":"ODC-ODbL","Link":"https........
}

func ExampleGetLastUpdate() {
	datetime, err := dawum.GetLastUpdate(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(datetime)
	// 2023-09-29T09:41:59+02:00
}
