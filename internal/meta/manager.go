package meta

import "github.com/orewaee/bytebin/pkg/dto"

type Manager interface {
	Load() error
	Unload() error
	Add(meta *dto.Meta) error
	RemoveById(id string) error
	GetById(id string) (*dto.Meta, error)
}
