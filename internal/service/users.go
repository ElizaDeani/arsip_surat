package service

import (
	"context"
	"errors"
	"time"

	"github.com/ElizaDeani/archivio/config"
	"github.com/ElizaDeani/archivio/internal/entity"
	"github.com/ElizaDeani/archivio/internal/http/dto"
	"github.com/ElizaDeani/archivio/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string, role string) (*entity.JWTCustomClaims, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	GetById(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, req dto.CreateUserRequest) error
	Update(ctx context.Context, req dto.UpdateUserRequest) error
	Delete(ctx context.Context, user *entity.User) error
}

type userService struct {
	cfg            *config.Config
	userRepository repository.UserRepository
}

func NewUserService(
	cfg *config.Config,
	userRepository repository.UserRepository,
) UserService {
	return &userService{cfg, userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string, role string) (*entity.JWTCustomClaims, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username atau password salah")
	}

	expiredTime := time.Now().Add(time.Hour * 10)

	claims := &entity.JWTCustomClaims{
		Username: user.Username,
		Role:     user.Role,
		Id:       int64(user.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "archivio",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	return claims, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     "user", // Default role untuk pengguna yang mendaftar
	}

	return s.userRepository.Create(ctx, user)
}

func (s *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}
func (s *userService) GetById(ctx context.Context, id int) (*entity.User, error) {
	return s.userRepository.GetById(ctx, id)
}

func (s *userService) Create(ctx context.Context, req dto.CreateUserRequest) error {
	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah digunakan")
	}

	user := &entity.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) error {
	user, err := s.userRepository.GetById(ctx, req.Id)
	if err != nil {
		return err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		user.Password = req.Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		return s.userRepository.Create(ctx, user)
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	return s.userRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, user *entity.User) error {
	return s.userRepository.Delete(ctx, user)
}
