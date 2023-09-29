package dto_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jxsl13/dawum/dto"
	"github.com/jxsl13/dawum/internal/testutils"
	"github.com/stretchr/testify/require"
)

var TestData map[string][]byte

func TestMain(m *testing.M) {
	root, files := testutils.FilePaths(`.*\.json`, "../testdata")
	TestData = make(map[string][]byte, len(files))

	for _, file := range files {
		absPath := filepath.Join(root, file)
		data, _ := os.ReadFile(absPath)
		TestData[file] = data
	}

	os.Exit(m.Run())
}

func TestDTO(t *testing.T) {
	data := TestData["basic.json"]

	var d dto.Data
	err := json.Unmarshal(data, &d)
	require.NoError(t, err)

	b, err := json.Marshal(d)
	require.NoError(t, err)

	var n dto.Data
	err = json.Unmarshal(b, &n)
	require.NoError(t, err)

	require.Equal(t, d, n)
}
