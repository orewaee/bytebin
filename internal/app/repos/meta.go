package repos

import "github.com/orewaee/bytebin/internal/app/domain"

type MetaRepo interface {
	AddMeta(id string, meta *domain.Meta) error
	RemoveMetaById(id string) error
	GetMetaById(id string) (*domain.Meta, error)
	GetAllMetaIds() ([]string, error)
}
