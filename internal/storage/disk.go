package storage

import (
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/meta"
	"github.com/orewaee/bytebin/pkg/dto"
)

type DiskStorage struct {
	bins  bin.Manager
	metas meta.Manager
}

func NewDiskStorage(bins bin.Manager, metas meta.Manager) *DiskStorage {
	return &DiskStorage{
		bins:  bins,
		metas: metas,
	}
}

func Load() error {
	panic("unimplemented")
}

func Unload() error {
	panic("unimplemented")
}

func Add(id string, bytes []byte, meta *dto.Meta) error {
	panic("unimplemented")
}

func RemoveById(id string) error {
	panic("unimplemented")
}

func GetById(id string) ([]byte, *dto.Meta, error) {
	panic("unimplemented")
}

func GetAllIds() ([]string, error) {
	panic("unimplemented")
}
