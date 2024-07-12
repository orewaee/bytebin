package storage

import "github.com/orewaee/bytebin/pkg/dto"

type Storage interface {
	Load() error
	Unload() error
	Add(id string, bytes []byte, meta *dto.Meta) error
	RemoveById(id string) error
	GetById(id string) ([]byte, *dto.Meta, error)
	GetAllIds() ([]string, error)
}
