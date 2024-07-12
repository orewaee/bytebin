package storage

import (
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/meta"
	"github.com/orewaee/bytebin/pkg/dto"
	"log"
	"slices"
	"time"
)

type DiskStorage struct {
	bins   bin.Manager
	metas  meta.Manager
	timers map[string]*time.Timer
}

func NewDiskStorage(bins bin.Manager, metas meta.Manager) *DiskStorage {
	return &DiskStorage{
		bins:   bins,
		metas:  metas,
		timers: make(map[string]*time.Timer),
	}
}

func (s *DiskStorage) Load() error {
	ids, err := s.GetAllIds()
	if err != nil {
		return err
	}

	for _, id := range ids {
		m, err := s.metas.GetById(id)
		if err != nil {
			return err
		}

		expireAt := m.CreatedAt.Add(m.Lifetime)
		if expireAt.After(time.Now()) {
			timer := time.AfterFunc(expireAt.Sub(time.Now()), func() {
				if err := s.RemoveById(id); err != nil {
					log.Println(err)
				}
			})

			s.timers[id] = timer

			continue
		}

		if err := s.RemoveById(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *DiskStorage) Unload() error {
	for _, timer := range s.timers {
		timer.Stop()
	}

	return nil
}

func (s *DiskStorage) Add(id string, bytes []byte, meta *dto.Meta) error {
	if err := s.bins.Add(id, bytes); err != nil {
		return err
	}

	if err := s.metas.Add(id, meta); err != nil {
		return err
	}

	timer := time.AfterFunc(meta.Lifetime, func() {
		if err := s.RemoveById(id); err != nil {
			log.Println(err)
		}
	})

	s.timers[id] = timer

	return nil
}

func (s *DiskStorage) RemoveById(id string) error {
	if err := s.bins.RemoveById(id); err != nil {
		return err
	}

	if err := s.metas.RemoveById(id); err != nil {
		return err
	}

	timer, ok := s.timers[id]
	if ok {
		timer.Stop()
	}

	return nil
}

func (s *DiskStorage) GetById(id string) ([]byte, *dto.Meta, error) {
	b, err := s.bins.GetById(id)
	if err != nil {
		return nil, nil, err
	}

	m, err := s.metas.GetById(id)
	if err != nil {
		return nil, nil, err
	}

	return b, m, nil
}

func (s *DiskStorage) GetAllIds() ([]string, error) {
	metaIds, err := s.metas.GetAllIds()
	if err != nil {
		return nil, err
	}

	binIds, err := s.bins.GetAllIds()
	if err != nil {
		return nil, err
	}

	var ids = make([]string, 0, len(metaIds))
	for _, id := range metaIds {
		if slices.Contains(binIds, id) {
			ids = append(ids, id)
		}
	}

	return ids, nil
}
