package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orewaee/bytebin/pkg/dto"
	"os"
	"strings"
)

type DiskManager struct{}

func NewDiskManager() *DiskManager {
	return &DiskManager{}
}

func (manager *DiskManager) Add(id string, meta *dto.Meta) error {
	path := fmt.Sprintf("./meta/meta-%s.json", id)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(meta)
}

func (manager *DiskManager) RemoveById(id string) error {
	path := fmt.Sprintf("./meta/meta-%s.json", id)
	return os.Remove(path)
}

func (manager *DiskManager) GetById(id string) (*dto.Meta, error) {
	path := fmt.Sprintf("./meta/meta-%s.json", id)

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

	var meta = new(dto.Meta)
	if err := json.Unmarshal(bytes, meta); err != nil {
		return nil, err
	}

	return meta, nil
}

func (manager *DiskManager) GetAllIds() ([]string, error) {
	path := "./meta"

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var ids = make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		id := strings.Replace(
			strings.TrimSuffix(entry.Name(), ".json"),
			"meta-", "", 1,
		)
		ids = append(ids, id)
	}

	return ids, nil
}
