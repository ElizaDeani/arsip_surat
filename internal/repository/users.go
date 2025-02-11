package repository

import (
	"context"

	"github.com/ElizaDeani/archivio/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetById(ctx context.Context, id int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	result := make([]entity.User, 0)

	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetById(ctx context.Context, id int) (*entity.User, error) {
	result := new(entity.User)

	if err := r.db.WithContext(ctx).Where("id_user = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Where("id_user = ?", user.Id).
		Updates(user).Error
}

func (r *userRepository) Delete(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Delete(&user).Error
}
