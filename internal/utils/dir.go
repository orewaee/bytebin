package utils

import (
	"errors"
	"fmt"
	"os"
)

// CheckDir creates a dir with the specified name if it does not exist.
// Returns an error if a file with the specified name exists.
func CheckDir(name string) error {
	info, err := os.Stat(name)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(name, 0700)
	}

	if err != nil {
		return err
	}

	if !info.IsDir() {
		text := fmt.Sprintf("%s is not a dir", name)
		return errors.New(text)
	}

	return nil
}
