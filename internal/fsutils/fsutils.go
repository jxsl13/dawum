package fsutils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type OsFS struct{}

func (OsFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func NewOsFS() OsFS {
	return OsFS{}
}

func ReadFile(fsys fs.FS, filePath string) ([]byte, error) {
	return fs.ReadFile(fsys, filePath)
}

func WalkDir(fsys fs.FS, dir string, walkFn fs.WalkDirFunc) error {
	return fs.WalkDir(fsys, dir, walkFn)
}

func FilePaths(fsys fs.FS, regex, rootDir string) (string, []string, error) {
	var files []string
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", nil, fmt.Errorf("cannot find file paths: invalid regex: %w", err)
	}

	prefix := filepath.Clean(rootDir) + string(filepath.Separator)
	err = WalkDir(fsys, rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !re.MatchString(path) {
			return nil
		}
		files = append(files, strings.TrimPrefix(path, prefix))
		return nil
	})

	if err != nil {
		return "", nil, fmt.Errorf("cannot find file paths: %w", err)
	}

	return rootDir, files, err
}
