package service

import (
	"context"

	"github.com/ElizaDeani/archivio/internal/entity"
	"github.com/ElizaDeani/archivio/internal/http/dto"
	"github.com/ElizaDeani/archivio/internal/repository"
)

type SuratKeluarService interface {
	GetAll(ctx context.Context) ([]entity.SuratKeluar, error)
	GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratKeluar, error)
	Create(ctx context.Context, req dto.CreateSuratKeluarRequest) error
	Update(ctx context.Context, req dto.UpdateSuratKeluarRequest) error
	Delete(ctx context.Context, suratKeluar *entity.SuratKeluar) error
}

type suratKeluarService struct {
	suratKeluarRepository repository.SuratKeluarRepository
}

func NewSuratKeluarService(suratKeluarRepository repository.SuratKeluarRepository) SuratKeluarService {
	return &suratKeluarService{suratKeluarRepository}
}

func (s *suratKeluarService) GetAll(ctx context.Context) ([]entity.SuratKeluar, error) {
	return s.suratKeluarRepository.GetAll(ctx)
}

func (s *suratKeluarService) GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratKeluar, error) {
	return s.suratKeluarRepository.GetByKodeSurat(ctx, kodeSurat)
}

func (s *suratKeluarService) Create(ctx context.Context, req dto.CreateSuratKeluarRequest) error {
	suratKeluar := &entity.SuratKeluar{
		KodeSurat:     req.KodeSurat,
		TanggalKeluar: req.TanggalKeluar,
		NoSurat:       req.NoSurat,
		TanggalSurat:  req.TanggalSurat,
		Pengirim:      req.Pengirim,
		Kepada:        req.Kepada,
		Perihal:       req.Perihal,
	}

	if err := s.suratKeluarRepository.Create(ctx, suratKeluar); err != nil {
		return err
	}
	return nil
}

func (s *suratKeluarService) Update(ctx context.Context, req dto.UpdateSuratKeluarRequest) error {
	suratKeluar, err := s.suratKeluarRepository.GetByKodeSurat(ctx, req.KodeSurat)
	if err != nil {
		return err
	}
	suratKeluar.KodeSurat = req.KodeSurat
	suratKeluar.TanggalKeluar = req.TanggalKeluar
	suratKeluar.NoSurat = req.NoSurat
	suratKeluar.TanggalSurat = req.TanggalSurat
	suratKeluar.Pengirim = req.Pengirim
	suratKeluar.Kepada = req.Kepada
	suratKeluar.Perihal = req.Perihal

	return s.suratKeluarRepository.Update(ctx, suratKeluar)
}

func (s *suratKeluarService) Delete(ctx context.Context, suratKeluar *entity.SuratKeluar) error {
	return s.suratKeluarRepository.Delete(ctx, suratKeluar)
}
