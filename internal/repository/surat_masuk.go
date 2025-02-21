package repository

import (
	"context"

	"github.com/ElizaDeani/archivio/internal/entity"
	"gorm.io/gorm"
)

type SuratMasukRepository interface {
	GetAll(ctx context.Context) ([]entity.SuratMasuk, error)
	GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratMasuk, error)
	Create(ctx context.Context, suratMasuk *entity.SuratMasuk) error
	Update(ctx context.Context, suratMasuk *entity.SuratMasuk) error
	Delete(ctx context.Context, suratMasuk *entity.SuratMasuk) error
	UpdateLampiran(ctx context.Context, kodeSurat int64, filePath string) error
}

type suratMasukRepository struct {
	db *gorm.DB
}

func NewSuratMasukRepository(db *gorm.DB) SuratMasukRepository {
	return &suratMasukRepository{db}
}

func (r *suratMasukRepository) GetAll(ctx context.Context) ([]entity.SuratMasuk, error) {
	result := make([]entity.SuratMasuk, 0)

	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *suratMasukRepository) GetByKodeSurat(ctx context.Context, kodeSurat int64) (*entity.SuratMasuk, error) {
	var suratMasuk entity.SuratMasuk
	err := r.db.WithContext(ctx).Where("kode_surat = ?", kodeSurat).First(&suratMasuk).Error
	if err != nil {
		return nil, err // Misalnya jika surat tidak ditemukan
	}
	return &suratMasuk, nil
}

func (r *suratMasukRepository) Create(ctx context.Context, suratMasuk *entity.SuratMasuk) error {
	return r.db.WithContext(ctx).Create(suratMasuk).Error
}

func (r *suratMasukRepository) Update(ctx context.Context, suratMasuk *entity.SuratMasuk) error {
	return r.db.WithContext(ctx).Save(suratMasuk).Error
}

func (r *suratMasukRepository) Delete(ctx context.Context, suratMasuk *entity.SuratMasuk) error {
	return r.db.WithContext(ctx).Delete(suratMasuk).Error
}

// UpdateLampiran hanya memperbarui kolom lampiran di database
func (r *suratMasukRepository) UpdateLampiran(ctx context.Context, kodeSurat int64, filePath string) error {
	return r.db.WithContext(ctx).Model(&entity.SuratMasuk{}).
		Where("kode_surat = ?", kodeSurat).
		Update("lampiran", filePath).Error
}
