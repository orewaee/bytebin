package services

import (
	"github.com/orewaee/bytebin/internal/app/domain"
	"github.com/orewaee/bytebin/internal/app/ports"
	"log"
	"slices"
	"time"
)

type BytebinService struct {
	binRepo  ports.BinRepo
	metaRepo ports.MetaRepo
	timers   map[string]*time.Timer
}

func NewBytebinService(binRepo ports.BinRepo, metaRepo ports.MetaRepo) *BytebinService {
	return &BytebinService{
		binRepo:  binRepo,
		metaRepo: metaRepo,
		timers:   make(map[string]*time.Timer),
	}
}

func (service *BytebinService) Load() error {
	ids, err := service.GetAllIds()
	if err != nil {
		return err
	}

	binIds, err := service.binRepo.GetAllIds()
	for _, binId := range binIds {
		if slices.Contains(ids, binId) {
			continue
		}

		if err := service.binRepo.RemoveById(binId); err != nil {
			return err
		}
	}

	metaIds, err := service.metaRepo.GetAllIds()
	for _, metaId := range metaIds {
		if slices.Contains(ids, metaId) {
			continue
		}

		if err := service.metaRepo.RemoveById(metaId); err != nil {
			return err
		}
	}

	for _, id := range ids {
		m, err := service.metaRepo.GetById(id)
		if err != nil {
			return err
		}

		expireAt := m.CreatedAt.Add(m.Lifetime)
		if expireAt.After(time.Now()) {

			duration := expireAt.Sub(time.Now())
			task := func() {
				if err := service.RemoveById(id); err != nil {
					log.Println(err)
				}
			}

			service.timers[id] = time.AfterFunc(duration, task)

			continue
		}

		if err := service.RemoveById(id); err != nil {
			return err
		}
	}

	return nil
}

func (service *BytebinService) Unload() error {
	for id, timer := range service.timers {
		timer.Stop()
		log.Println("bin", id, "timer stopped")
	}

	return nil
}

func (service *BytebinService) Add(id string, bin []byte, meta *domain.Meta) error {
	if err := service.binRepo.Add(id, bin); err != nil {
		return err
	}

	if err := service.metaRepo.Add(id, meta); err != nil {
		return err
	}

	task := func() {
		if err := service.RemoveById(id); err != nil {
			log.Println(err) // ??
		}
	}

	service.timers[id] = time.AfterFunc(meta.Lifetime, task)

	return nil
}

func (service *BytebinService) RemoveById(id string) error {
	if err := service.binRepo.RemoveById(id); err != nil {
		return err
	}

	if err := service.metaRepo.RemoveById(id); err != nil {
		return err
	}

	timer, ok := service.timers[id]
	if ok {
		timer.Stop()
	}

	return nil
}

func (service *BytebinService) GetById(id string) ([]byte, *domain.Meta, error) {
	bin, err := service.binRepo.GetById(id)
	if err != nil {
		return nil, nil, err
	}

	meta, err := service.metaRepo.GetById(id)
	if err != nil {
		return nil, nil, err
	}

	return bin, meta, nil
}

func (service *BytebinService) GetAllIds() ([]string, error) {
	binIds, err := service.binRepo.GetAllIds()
	if err != nil {
		return nil, err
	}

	metaIds, err := service.metaRepo.GetAllIds()
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
