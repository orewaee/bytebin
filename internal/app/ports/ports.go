package ports

import "github.com/orewaee/bytebin/internal/app/domain"

type BinRepo interface {
	Add(id string, bin []byte) error
	RemoveById(id string) error
	GetById(id string) ([]byte, error)
	GetAllIds() ([]string, error)
}

type MetaRepo interface {
	Add(id string, meta *domain.Meta) error
	RemoveById(id string) error
	GetById(id string) (*domain.Meta, error)
	GetAllIds() ([]string, error)
}

type BytebinService interface {
	Load() error
	Unload() error
	Add(id string, bin []byte, meta *domain.Meta) error
	RemoveById(id string) error
	GetById(id string) ([]byte, *domain.Meta, error)
	GetAllIds() ([]string, error)
}
