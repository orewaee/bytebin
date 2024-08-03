package bin

import (
	"errors"
	"os"
	"strings"
)

type DiskBinRepo struct{}

func NewDiskBinRepo() *DiskBinRepo {
	return &DiskBinRepo{}
}

func (repo *DiskBinRepo) Add(id string, bin []byte) error {
	path := "./bins/bin-" + id

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := file.Write(bin); err != nil {
		return err
	}

	return file.Close()
}

func (repo *DiskBinRepo) RemoveById(id string) error {
	path := "./bins/bin-" + id
	return os.Remove(path)
}

func (repo *DiskBinRepo) GetById(id string) ([]byte, error) {
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

func (repo *DiskBinRepo) GetAllIds() ([]string, error) {
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
