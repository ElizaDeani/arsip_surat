package service

import (
	"context"

	"github.com/ElizaDeani/archivio/internal/entity"
	"github.com/ElizaDeani/archivio/internal/http/dto"
	"github.com/ElizaDeani/archivio/internal/repository"
)

type SuratMasukService interface {
	GetAll(ctx context.Context) ([]entity.SuratMasuk, error)
	GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratMasuk, error)
	Create(ctx context.Context, req dto.CreateSuratMasukRequest) error
	Update(ctx context.Context, req dto.UpdateSuratMasukRequest) error
	Delete(ctx context.Context, suratMasuk *entity.SuratMasuk) error
}

type suratMasukService struct {
	suratMasukRepository repository.SuratMasukRepository
}

func NewSuratMasukService(suratMasukRepository repository.SuratMasukRepository) SuratMasukService {
	return &suratMasukService{suratMasukRepository}
}

func (s *suratMasukService) GetAll(ctx context.Context) ([]entity.SuratMasuk, error) {
	return s.suratMasukRepository.GetAll(ctx)
}

func (s *suratMasukService) GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratMasuk, error) {
	return s.suratMasukRepository.GetByKodeSurat(ctx, kodeSurat)
}

func (s *suratMasukService) Create(ctx context.Context, req dto.CreateSuratMasukRequest) error {
	// Buat entitas SuratMasuk dari request DTO
	suratMasuk := &entity.SuratMasuk{
		KodeSurat:    req.KodeSurat,
		TanggalMasuk: req.TanggalMasuk,
		NoSurat:      req.NoSurat,
		TanggalSurat: req.TanggalSurat,
		Pengirim:     req.Pengirim,
		Kepada:       req.Kepada,
		Perihal:      req.Perihal,
	}
	// Simpan ke repository
	if err := s.suratMasukRepository.Create(ctx, suratMasuk); err != nil {
		return err
	}
	return nil
}

// Update implements TicketService.
func (s *suratMasukService) Update(ctx context.Context, req dto.UpdateSuratMasukRequest) error {
	suratMasuk, err := s.suratMasukRepository.GetByKodeSurat(ctx, req.KodeSurat)
	if err != nil {
		return err
	}
	suratMasuk.KodeSurat = req.KodeSurat
	suratMasuk.TanggalMasuk = req.TanggalMasuk
	suratMasuk.NoSurat = req.NoSurat
	suratMasuk.TanggalSurat = req.TanggalSurat
	suratMasuk.Pengirim = req.Pengirim
	suratMasuk.Kepada = req.Kepada
	suratMasuk.Perihal = req.Perihal

	return s.suratMasukRepository.Update(ctx, suratMasuk)
}

// Delete implements TicketService.
func (s *suratMasukService) Delete(ctx context.Context, suratMasuk *entity.SuratMasuk) error {
	return s.suratMasukRepository.Delete(ctx, suratMasuk)

}
