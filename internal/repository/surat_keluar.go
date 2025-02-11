package repository

import (
	"context"

	"github.com/ElizaDeani/archivio/internal/entity"
	"gorm.io/gorm"
)

type SuratKeluarRepository interface {
	GetAll(ctx context.Context) ([]entity.SuratKeluar, error)
	GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratKeluar, error)
	Create(ctx context.Context, suratKeluar *entity.SuratKeluar) error
	Update(ctx context.Context, suratkeluar *entity.SuratKeluar) error
	Delete(ctx context.Context, suratKeluar *entity.SuratKeluar) error
}

type suratKeluarRepository struct {
	db *gorm.DB
}

func NewSuratKeluarRepository(db *gorm.DB) SuratKeluarRepository {
	return &suratKeluarRepository{db}
}

func (r *suratKeluarRepository) GetAll(ctx context.Context) ([]entity.SuratKeluar, error) {
	result := make([]entity.SuratKeluar, 0)

	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, nil
	}
	return result, nil
}

func (r *suratKeluarRepository) GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratKeluar, error) {
	var suratKeluar entity.SuratKeluar
	err := r.db.WithContext(ctx).Where("kode_surat = ?", kodeSurat).First(&suratKeluar).Error
	if err != nil {
		return nil, err // Misalnya jika surat tidak ditemukan
	}
	return &suratKeluar, nil
}

func (r *suratKeluarRepository) Create(ctx context.Context, suratKeluar *entity.SuratKeluar) error {
	return r.db.WithContext(ctx).Create(suratKeluar).Error
}

func (r *suratKeluarRepository) Update(ctx context.Context, suratKeluar *entity.SuratKeluar) error {
	return r.db.WithContext(ctx).Save(suratKeluar).Error
}

func (r *suratKeluarRepository) Delete(ctx context.Context, suratKeluar *entity.SuratKeluar) error {
	return r.db.WithContext(ctx).Delete(suratKeluar).Error
}
