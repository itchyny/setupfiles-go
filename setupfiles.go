package setupfiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Create the files.
func Create(dir string, source string) error {
	files, err := parse(source)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(dir, f.path)
		if !strings.HasPrefix(path, dir) {
			return errors.Errorf("invalid path: %s", f.path)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		if f.symlink != "" {
			link := filepath.Join(filepath.Dir(path), f.symlink)
			if !strings.HasPrefix(link, dir) {
				return errors.Errorf("invalid path: %s", f.symlink)
			}
			if err := os.Symlink(link, path); err != nil {
				return err
			}
			continue
		}
		if f.isDir {
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
			continue
		}
		if err := ioutil.WriteFile(path, []byte(f.contents), 0644); err != nil {
			return err
		}
	}
	return nil
}

// CreateTemp creates the files in a temporary directory with the given prefix.
func CreateTemp(prefix string, source string) (string, error) {
	root, err := ioutil.TempDir("", prefix)
	if err != nil {
		return "", err
	}
	if err = Create(root, source); err != nil {
		_ = os.Remove(root)
		return "", err
	}
	return root, nil
}
