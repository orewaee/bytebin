package services

import (
	"github.com/orewaee/bytebin/internal/app/domain"
	"github.com/orewaee/bytebin/internal/app/repos"
	"github.com/rs/zerolog"
	"slices"
	"time"
)

type BytebinService struct {
	binRepo  repos.BinRepo
	metaRepo repos.MetaRepo
	timers   map[string]*time.Timer
	log      *zerolog.Logger
}

func NewBytebinService(binRepo repos.BinRepo, metaRepo repos.MetaRepo, log *zerolog.Logger) *BytebinService {
	return &BytebinService{
		log:      log,
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

	binIds, err := service.binRepo.GetAllBinIds()
	for _, binId := range binIds {
		if slices.Contains(ids, binId) {
			continue
		}

		if err := service.binRepo.RemoveBinById(binId); err != nil {
			return err
		}
	}

	metaIds, err := service.metaRepo.GetAllMetaIds()
	for _, metaId := range metaIds {
		if slices.Contains(ids, metaId) {
			continue
		}

		if err := service.metaRepo.RemoveMetaById(metaId); err != nil {
			service.log.Err(err).Send()
			return err
		}
	}

	for _, id := range ids {
		m, err := service.metaRepo.GetMetaById(id)
		if err != nil {
			service.log.Err(err).Send()
			return err
		}

		expireAt := m.CreatedAt.Add(m.Lifetime)
		if expireAt.After(time.Now()) {
			duration := expireAt.Sub(time.Now())
			task := func() {
				if err := service.RemoveById(id); err != nil {
					service.log.Err(err).Send()
				}
			}

			service.timers[id] = time.AfterFunc(duration, task)

			service.log.Debug().
				Str("id", id).
				Msg("bin and meta loaded")

			continue
		}

		if err := service.RemoveById(id); err != nil {
			service.log.Err(err).Send()
			return err
		}
	}

	return nil
}

func (service *BytebinService) Unload() error {
	for id, timer := range service.timers {
		timer.Stop()

		service.log.Debug().
			Str("id", id).
			Msg("timer stopped")
	}

	return nil
}

func (service *BytebinService) Add(id string, bin []byte, meta *domain.Meta) error {
	if err := service.binRepo.AddBin(id, bin); err != nil {
		service.log.Err(err).Send()
		return err
	}

	if err := service.metaRepo.AddMeta(id, meta); err != nil {
		service.log.Err(err).Send()
		return err
	}

	task := func() {
		if err := service.RemoveById(id); err != nil {
			service.log.Err(err).Send()
		}
	}

	service.timers[id] = time.AfterFunc(meta.Lifetime, task)

	service.log.Info().
		Str("id", id).
		Str("content_type", meta.ContentType).
		Str("user_agent", meta.UserAgent).
		Msg("bin and meta added")

	return nil
}

func (service *BytebinService) RemoveById(id string) error {
	if err := service.binRepo.RemoveBinById(id); err != nil {
		service.log.Err(err).Send()
		return err
	}

	if err := service.metaRepo.RemoveMetaById(id); err != nil {
		service.log.Err(err).Send()
		return err
	}

	timer, ok := service.timers[id]
	if ok {
		timer.Stop()

		service.log.Debug().
			Str("id", id).
			Msg("timer stopped")
	}

	return nil
}

func (service *BytebinService) GetById(id string) ([]byte, *domain.Meta, error) {
	bin, err := service.binRepo.GetBinById(id)
	if err != nil {
		service.log.Err(err).Send()
		return nil, nil, err
	}

	meta, err := service.metaRepo.GetMetaById(id)
	if err != nil {
		service.log.Err(err).Send()
		return nil, nil, err
	}

	return bin, meta, nil
}

func (service *BytebinService) GetAllIds() ([]string, error) {
	binIds, err := service.binRepo.GetAllBinIds()
	if err != nil {
		service.log.Err(err).Send()
		return nil, err
	}

	metaIds, err := service.metaRepo.GetAllMetaIds()
	if err != nil {
		service.log.Err(err).Send()
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
