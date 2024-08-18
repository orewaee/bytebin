package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orewaee/bytebin/internal/app/domain"
	"os"
	"strings"
)

type DiskMetaRepo struct{}

func NewDiskMetaRepo() *DiskMetaRepo {
	return &DiskMetaRepo{}
}

func (repo *DiskMetaRepo) AddMeta(id string, meta *domain.Meta) error {
	path := fmt.Sprintf("./metas/meta-%s.json", id)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	data := &Meta{
		Id:          meta.Id,
		ContentType: meta.ContentType,
		Ip:          meta.Ip,
		UserAgent:   meta.UserAgent,
		CreatedAt:   meta.CreatedAt,
		Lifetime:    meta.Lifetime,
	}

	if err := json.NewEncoder(file).Encode(data); err != nil {
		return err
	}

	return file.Close()
}

func (repo *DiskMetaRepo) RemoveMetaById(id string) error {
	path := fmt.Sprintf("./metas/meta-%s.json", id)
	return os.Remove(path)
}

func (repo *DiskMetaRepo) GetMetaById(id string) (*domain.Meta, error) {
	path := fmt.Sprintf("./metas/meta-%s.json", id)

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

	var data = new(Meta)
	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	meta := &domain.Meta{
		Id:          data.Id,
		ContentType: data.ContentType,
		Ip:          data.Ip,
		UserAgent:   data.UserAgent,
		CreatedAt:   data.CreatedAt,
		Lifetime:    data.Lifetime,
	}

	return meta, nil
}

func (repo *DiskMetaRepo) GetAllMetaIds() ([]string, error) {
	path := "./metas"

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
