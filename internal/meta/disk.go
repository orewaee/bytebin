package meta

import "github.com/orewaee/bytebin/pkg/dto"

type DiskManager struct{}

func NewDiskManager() *DiskManager {
	return &DiskManager{}
}

func (manager *DiskManager) Load() error {
	panic("unimplemented")
}

func (manager *DiskManager) Unload() error {
	panic("unimplemented")
}

func (manager *DiskManager) Add(meta *dto.Meta) error {
	panic("unimplemented")
}

func (manager *DiskManager) RemoveById(id string) error {
	panic("unimplemented")
}

func (manager *DiskManager) GetById(id string) (*dto.Meta, error) {
	panic("unimplemented")
}
