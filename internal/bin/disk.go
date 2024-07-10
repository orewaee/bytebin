package bin

import (
	"errors"
	"os"
	"strings"
)

type DiskManager struct{}

func NewDiskManager() *DiskManager {
	return &DiskManager{}
}

func (manager *DiskManager) Add(id string, bytes []byte) error {
	path := "./bins/bin-" + id

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := file.Write(bytes); err != nil {
		return err
	}

	return nil
}

func (manager *DiskManager) RemoveById(id string) error {
	path := "./bins/bin-" + id
	return os.Remove(path)
}

func (manager *DiskManager) GetById(id string) ([]byte, error) {
	path := "./bins/bin-" + id

	info, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, ErrNotExist
	}

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, ErrInvalid
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (manager *DiskManager) GetAllIds() ([]string, error) {
	path := "./bins"

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var ids = make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		id := strings.TrimPrefix(entry.Name(), "bin-")
		ids = append(ids, id)
	}

	return ids, nil
}
