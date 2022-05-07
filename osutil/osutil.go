package osutil

import (
	"errors"
	"io/fs"
	"os"
)

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return err == nil, err
}
