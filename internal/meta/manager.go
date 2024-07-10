package meta

import "github.com/orewaee/bytebin/pkg/dto"

type Manager interface {
	Add(id string, meta *dto.Meta) error
	RemoveById(id string) error
	GetById(id string) (*dto.Meta, error)
	GetAllIds() ([]string, error)
}
