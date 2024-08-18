package api

import "github.com/orewaee/bytebin/internal/app/domain"

type BytebinApi interface {
	Load() error
	Unload() error
	Add(id string, bin []byte, meta *domain.Meta) error
	RemoveById(id string) error
	GetById(id string) ([]byte, *domain.Meta, error)
	GetAllIds() ([]string, error)
}
